package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"

	pb "github.com/Chocobone/GDG_backend_sideProject/proto"
	"google.golang.org/grpc"
)

// MarketDataServiceServer 구현
type server struct {
	pb.UnimplementedMarketDataServiceServer
}

// REST API에서 데이터를 가져오는 함수
func fetchMarketData(market string) (*pb.MarketDataResponse, error) {
	apiURL := fmt.Sprintf("https://api.upbit.com/v1/ticker?markets=%s", market)
	resp, err := http.Get(apiURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %v", err)
	}
	defer resp.Body.Close()

	var apiResponse []struct {
		Market           string  `json:"market"`
		TradeDate        string  `json:"trade_date"`
		TradeTime        string  `json:"trade_time"`
		TradeTimestamp   int64   `json:"trade_timestamp"`
		OpeningPrice     float64 `json:"opening_price"`
		HighPrice        float64 `json:"high_price"`
		LowPrice         float64 `json:"low_price"`
		TradePrice       float64 `json:"trade_price"`
		PrevClosingPrice float64 `json:"prev_closing_price"`
		ChangePrice      float64 `json:"change_price"`
		ChangeRate       float64 `json:"change_rate"`
		TradeVolume      float64 `json:"trade_volume"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}

	if len(apiResponse) == 0 {
		return nil, fmt.Errorf("no data received")
	}

	data := apiResponse[0]
	return &pb.MarketDataResponse{
		Market:           data.Market,
		TradeDate:        data.TradeDate,
		TradeTime:        data.TradeTime,
		TradeTimestamp:   data.TradeTimestamp,
		OpeningPrice:     data.OpeningPrice,
		HighPrice:        data.HighPrice,
		LowPrice:         data.LowPrice,
		TradePrice:       data.TradePrice,
		PrevClosingPrice: data.PrevClosingPrice,
		ChangePrice:      data.ChangePrice,
		ChangeRate:       data.ChangeRate,
		TradeVolume:      data.TradeVolume,
	}, nil
}

// GetMarketData 메서드 구현
func (s *server) GetMarketData(ctx context.Context, req *pb.MarketDataRequest) (*pb.MarketDataResponse, error) {
	marketData, err := fetchMarketData(req.Market)
	if err != nil {
		return nil, err
	}
	return marketData, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterMarketDataServiceServer(grpcServer, &server{})

	log.Println("gRPC server is running on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
