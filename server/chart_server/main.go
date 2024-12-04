package main

import (
	"context"
	pb "github.com/Chocobone/GDG_backend_sideProject/server/pb"
	"log"
	"net/http"
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"google.golang.org/grpc"
)

// 실시간 데이터를 저장할 전역 변수
var candleData = [][4]int64{}
var timestamps = []string{}

// gRPC로 데이터 가져오기
func fetchBitcoinData() {
	// gRPC 서버에 연결
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("gRPC 서버 연결 실패: %v", err)
	}
	defer conn.Close()

	client := pb.NewMonitoring_ServiceClient(conn)
	req := &pb.MonitoringRequest{StartMonitoring: "KRW-BTC"}

	// SayHello 호출
	stream, err := client.SayHello(context.Background(), req)
	if err != nil {
		log.Fatalf("데이터 스트림 열기 실패: %v", err)
	}

	for {
		resp, err := stream.Recv()
		if err != nil {
			log.Printf("데이터 수신 중 오류: %v", err)
			return
		}

		// 데이터 저장
		candleData = append(candleData, [4]int64{resp.OpeningPrice, resp.HighPrice, resp.LowPrice, resp.PrevClosingPrice})
		timestamps = append(timestamps, time.Unix(resp.Timestamp, 0).Format("15:04:05"))

		// 로그 출력
		log.Printf("수신한 데이터: Open: %.2f, High: %.2f, Low: %.2f, Close: %.2f, Timestamp: %s",
			resp.OpeningPrice, resp.HighPrice, resp.LowPrice, resp.PrevClosingPrice, time.Unix(resp.Timestamp, 0).Format("15:04:05"))

		// 최근 100개의 데이터만 유지
		if len(candleData) > 100 {
			candleData = candleData[1:]
			timestamps = timestamps[1:]
		}
	}
}

// 캔들 차트 생성
func generateKLineChart() *charts.Kline {
	kline := charts.NewKLine()

	kline.SetGlobalOptions(
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
			Name: "시간",
			Type: "category",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name:      "비트코인 가격 (KRW)",
			Type:      "value",
			Min:       "dataMin", // 데이터의 최소값으로 자동 설정
			AxisLabel: &opts.AxisLabel{Formatter: "{value} 원"},
		}),
	)

	// 캔들 데이터 추가
	klineData := make([]opts.KlineData, len(candleData))
	for i, v := range candleData {
		klineData[i] = opts.KlineData{Value: [4]int64{v[0], v[1], v[2], v[3]}}
	}

	kline.SetXAxis(timestamps).AddSeries("BTC 가격", klineData)

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
