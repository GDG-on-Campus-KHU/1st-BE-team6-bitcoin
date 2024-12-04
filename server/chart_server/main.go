package main

import (
	"context"
	pb "github.com/Chocobone/GDG_backend_sideProject/server/pb"
	"log"
	"net/http"
	"sort"
	"sync"
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"google.golang.org/grpc"
)

// 데이터 저장 및 동기화를 위한 구조체
type ChartData struct {
	mu         sync.Mutex // 동시 접근을 막기 위한 Mutex
	candleData [][4]int64 // 캔들 데이터 (시작가, 최고가, 최저가, 종가)
	timestamps []string   // 시간 데이터 (x축 라벨)
}

// chartData 전역 변수로 선언
var chartData = &ChartData{}

// TradeData는 거래 데이터를 나타냅니다.
type TradeData struct {
	Timestamp  int64 // Unix Timestamp (milliseconds)
	TradePrice int64 // 현재 거래가
}

// GroupedData는 1분 단위로 묶인 데이터를 나타냅니다.
type GroupedData struct {
	Minute       string // 1분 단위 시간 (HH:mm)
	OpeningPrice int64  // 시작가 (첫 번째 거래)
	HighPrice    int64  // 최고가
	LowPrice     int64  // 최저가
	ClosingPrice int64  // 종가 (마지막 거래)
}

// 1분 단위로 데이터를 그룹화하고 시작가, 최고가, 최저가, 종가를 계산하는 함수
func calculateMinutePrices(data []TradeData) []GroupedData {
	// 데이터를 1분 단위로 그룹화
	grouped := make(map[string][]TradeData)
	for _, trade := range data {
		// Unix Timestamp를 1분 단위 시간 문자열로 변환
		minute := time.Unix(trade.Timestamp/1000, 0).Format("15:04")
		grouped[minute] = append(grouped[minute], trade)
	}

	// 결과 저장
	var result []GroupedData
	for minute, trades := range grouped {
		// 시간 순으로 정렬
		sort.Slice(trades, func(i, j int) bool {
			return trades[i].Timestamp < trades[j].Timestamp
		})

		// 최고가 및 최저가 초기값 설정
		highPrice := trades[0].TradePrice
		lowPrice := trades[0].TradePrice

		// 최고가와 최저가 계산
		for _, trade := range trades {
			// 로그 출력으로 계산 과정 확인
			log.Printf("Minute: %s, TradePrice: %.2f", minute, trade.TradePrice)
			if trade.TradePrice > highPrice {
				highPrice = trade.TradePrice
			}
			if trade.TradePrice < lowPrice {
				lowPrice = trade.TradePrice
			}
		}

		// 결과 추가
		result = append(result, GroupedData{
			Minute:       minute,
			OpeningPrice: trades[0].TradePrice,             // 첫 번째 거래가
			HighPrice:    highPrice,                        // 최고가
			LowPrice:     lowPrice,                         // 최저가
			ClosingPrice: trades[len(trades)-1].TradePrice, // 마지막 거래가
		})
	}

	// 시간 순으로 결과 정렬
	sort.Slice(result, func(i, j int) bool {
		return result[i].Minute < result[j].Minute
	})

	return result
}

// gRPC로 데이터 가져오기
func fetchBitcoinData() {
	// gRPC 서버 연결
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("gRPC 서버 연결 실패: %v", err)
	}
	defer conn.Close()

	client := pb.NewMonitoring_ServiceClient(conn)
	req := &pb.MonitoringRequest{StartMonitoring: "KRW-BTC"}

	stream, err := client.SayHello(context.Background(), req)
	if err != nil {
		log.Fatalf("데이터 스트림 열기 실패: %v", err)
	}

	// 실시간 거래 데이터 저장소
	var rawData []TradeData

	for {
		// 데이터 수신
		resp, err := stream.Recv()
		if err != nil {
			log.Printf("데이터 수신 중 오류: %v", err)
			return
		}

		// 데이터 로그 출력
		log.Printf("수신 데이터 - Timestamp: %d, TradePrice: %.2f, HighPrice: %.2f, LowPrice: %.2f",
			resp.Timestamp, resp.TradePrice, resp.HighPrice, resp.LowPrice)

		// 거래 데이터를 rawData에 추가
		rawData = append(rawData, TradeData{
			Timestamp:  resp.Timestamp,
			TradePrice: resp.TradePrice,
		})

		// 1분 단위로 데이터를 그룹화
		groupedData := calculateMinutePrices(rawData)

		// 차트 데이터 갱신
		chartData.mu.Lock()
		chartData.candleData = nil
		chartData.timestamps = nil

		for _, group := range groupedData {
			chartData.candleData = append(chartData.candleData, [4]int64{
				int64(group.OpeningPrice), // 시작가
				int64(group.HighPrice),    // 최고가
				int64(group.LowPrice),     // 최저가
				int64(group.ClosingPrice), // 종가
			})
			chartData.timestamps = append(chartData.timestamps, group.Minute)
		}

		// 최근 100개의 데이터만 유지
		if len(chartData.candleData) > 100 {
			chartData.candleData = chartData.candleData[len(chartData.candleData)-100:]
			chartData.timestamps = chartData.timestamps[len(chartData.timestamps)-100:]
		}
		chartData.mu.Unlock()
	}
}

// 캔들 차트 생성
func generateKLineChart() *charts.Kline {
	kline := charts.NewKLine()

	chartData.mu.Lock()
	defer chartData.mu.Unlock()

	if len(chartData.candleData) == 0 {
		// 데이터가 없는 경우 빈 차트를 반환
		return kline
	}

	// 캔들 데이터 추가
	klineData := make([]opts.KlineData, len(chartData.candleData))
	minPrice := chartData.candleData[0][2] // 초기값: 첫 번째 데이터의 최저가
	maxPrice := chartData.candleData[0][1] // 초기값: 첫 번째 데이터의 최고가

	for i, v := range chartData.candleData {
		klineData[i] = opts.KlineData{Value: [4]int64{v[0], v[1], v[2], v[3]}}
		// 로그로 데이터 확인
		log.Printf("캔들 데이터 - Open: %d, High: %d, Low: %d, Close: %d", v[0], v[1], v[2], v[3])

		// Y축 범위 계산
		if v[1] > maxPrice {
			maxPrice = v[1]
		}
		if v[2] < minPrice {
			minPrice = v[2]
		}
	}

	// Y축 마진 설정 (10% 여유)
	margin := (maxPrice - minPrice) / 10
	yMin := float64(minPrice - margin)
	yMax := float64(maxPrice + margin)

	// 차트 설정
	kline.SetXAxis(chartData.timestamps).
		AddSeries("BTC 가격", klineData).
		SetGlobalOptions(
			charts.WithTitleOpts(opts.Title{
				Title:    "실시간 비트코인 시세 (캔들 차트)",
				Subtitle: "KRW-BTC 기준",
				Left:     "center",
			}),
			charts.WithTooltipOpts(opts.Tooltip{
				Show:    opts.Bool(true),
				Trigger: "axis",
			}),
			charts.WithXAxisOpts(opts.XAxis{
				Name:      "시간",
				Type:      "category",
				AxisLabel: &opts.AxisLabel{Rotate: 45}, // X축 라벨 회전
			}),
			charts.WithYAxisOpts(opts.YAxis{
				Name:      "비트코인 가격 (KRW)",
				Type:      "value",
				Min:       yMin, // 동적 Y축 최소값
				Max:       yMax, // 동적 Y축 최대값
				AxisLabel: &opts.AxisLabel{Formatter: "{value} 원"},
			}),
			charts.WithGridOpts(opts.Grid{
				Left:   "10%",
				Right:  "10%",
				Bottom: "15%",
			}),
		)

	return kline
}

// HTTP 핸들러
func handler(w http.ResponseWriter, r *http.Request) {
	chart := generateKLineChart()
	chart.Render(w)
}

func main() {
	// 데이터 수신 고루틴 실행
	go fetchBitcoinData()

	// HTTP 서버 설정
	http.HandleFunc("/", handler)
	log.Println("chart_server가 8080 포트에서 실행 중...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
