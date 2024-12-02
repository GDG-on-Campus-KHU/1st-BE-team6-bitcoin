package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	pb "github.com/Chocobone/GDG_backend_sideProject/server/pb" // proto 파일에서 생성된 패키지 경로)
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedMonitoring_ServiceServer
}

// REST API에서 데이터를 주기적으로 가져오는 함수
func fetchMarketData(market string) (*pb.MonitoringResponse, error) {
	apiURL := fmt.Sprintf("https://api.upbit.com/v1/ticker?markets=%s", market)
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %v", err)
	}
	defer resp.Body.Close()

	var apiResponse []struct {
		Market             string  `json:"market"`
		TradeDate          string  `json:"trade_date"`
		TradeTime          string  `json:"trade_time"`
		TradeDateKst       string  `json:"trade_date_kst"`
		TradeTimeKst       string  `json:"trade_time_kst"`
		TradeTimestamp     int64   `json:"trade_timestamp"`
		OpeningPrice       float64 `json:"opening_price"`
		HighPrice          float64 `json:"high_price"`
		LowPrice           float64 `json:"low_price"`
		TradePrice         float64 `json:"trade_price"`
		PrevClosingPrice   float64 `json:"prev_closing_price"`
		Change             string  `json:"change"`
		ChangePrice        float64 `json:"change_price"`
		ChangeRate         float64 `json:"change_rate"`
		SignedChangePrice  float64 `json:"signed_change_price"`
		SignedChangeRate   float64 `json:"signed_change_rate"`
		TradeVolume        float64 `json:"trade_volume"`
		AccTradePrice      float64 `json:"acc_trade_price"`
		AccTradePrice24h   float64 `json:"acc_trade_price_24h"`
		AccTradeVolume     float64 `json:"acc_trade_volume"`
		AccTradeVolume24h  float64 `json:"acc_trade_volume_24h"`
		Highest52WeekPrice float64 `json:"highest_52_week_price"`
		Highest52WeekDate  string  `json:"highest_52_week_date"`
		Lowest52WeekPrice  float64 `json:"lowest_52_week_price"`
		Lowest52WeekDate   string  `json:"lowest_52_week_date"`
		Timestamp          int64   `json:"timestamp"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	if len(apiResponse) == 0 {
		return nil, fmt.Errorf("no data received")
	}

	data := apiResponse[0]
	return &pb.MonitoringResponse{
		Market:              data.Market,
		TradeDate:           data.TradeDate,
		TradeTime:           data.TradeTime,
		TradeDateKst:        data.TradeDateKst,
		TradeTimeKst:        data.TradeTimeKst,
		TradeTimestamp:      data.TradeTimestamp,
		OpeningPrice:        int64(data.OpeningPrice),
		HighPrice:           int64(data.HighPrice),
		LowPrice:            int64(data.LowPrice),
		TradePrice:          int64(data.TradePrice),
		PrevClosingPrice:    int64(data.PrevClosingPrice),
		Change:              data.Change,
		ChangePrice:         int64(data.ChangePrice),
		ChangeRate:          data.ChangeRate,
		SignedChangePrice:   int64(data.SignedChangePrice),
		SignedChangeRate:    data.SignedChangeRate,
		TradeVolume:         data.TradeVolume,
		AccTradePrice:       data.AccTradePrice,
		AccTradePrice_24H:   data.AccTradePrice24h,
		AccTradeVolume:      data.AccTradeVolume,
		AccTradeVolume_24H:  data.AccTradeVolume24h,
		Highest_52WeekPrice: int64(data.Highest52WeekPrice),
		Highest_52WeekDate:  data.Highest52WeekDate,
		Lowest_52WeekPrice:  int64(data.Lowest52WeekPrice),
		Lowest_52WeekDate:   data.Lowest52WeekDate,
		Timestamp:           data.Timestamp,
	}, nil
}

func (s *server) SayHello(req *pb.MonitoringRequest, stream pb.Monitoring_Service_SayHelloServer) error {
	for {
		data, err := fetchMarketData(req.StartMonitoring)
		if err != nil {
			return err
		}

		if err := stream.Send(data); err != nil {
			return err
		}
		time.Sleep(1 * time.Second) // 5초마다 데이터를 스트리밍
	}
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterMonitoring_ServiceServer(grpcServer, &server{})

	log.Println("gRPC server running on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
