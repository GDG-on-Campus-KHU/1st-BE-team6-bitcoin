package chart

import (
	"context"
	"github.com/Chocobone/GDG_backend_sideProject/server/chart_server/util"
	"github.com/Chocobone/GDG_backend_sideProject/server/pb"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"google.golang.org/grpc"
	"log"
	"sync"
)

// 데이터 저장 및 동기화를 위한 구조체
type ChartData struct {
	mu         sync.Mutex // 동시 접근을 막기 위한 Mutex
	candleData [][4]int64 // 캔들 데이터 (시작가, 최고가, 최저가, 종가)
	timestamps []string   // 시간 데이터 (x축 라벨)
}

// chartData 전역 변수로 선언
var chartData = &ChartData{}

// 캔들 차트 생성
func GenerateKLineChart() *charts.Kline {
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

// gRPC로 데이터 가져오기
func FetchBitcoinData() {
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
	var rawData []util.TradeData

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
		rawData = append(rawData, util.TradeData{
			Timestamp:  resp.Timestamp,
			TradePrice: resp.TradePrice,
		})

		// 1분 단위로 데이터를 그룹화
		groupedData := util.CalculateMinutePrices(rawData)

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
