package grpc

import (
	"context"
	"github.com/Chocobone/GDG_backend_sideProject/server/chart_server/model"
	"github.com/Chocobone/GDG_backend_sideProject/server/chart_server/util"
	"github.com/Chocobone/GDG_backend_sideProject/server/pb"
	"google.golang.org/grpc"
	"log"
)

var tradeData []util.TradeData // 실시간 거래 데이터 저장

// gRPC 데이터 수신
func StartFetchingData() {
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

	for {
		resp, err := stream.Recv()
		if err != nil {
			log.Printf("데이터 수신 중 오류: %v", err)
			return
		}

		// 새로운 데이터 추가
		tradeData = append(tradeData, util.TradeData{
			Timestamp:  resp.Timestamp,
			TradePrice: int64(resp.TradePrice),
		})

		// 1분 단위 데이터 계산 후 모델 업데이트
		groupedData := util.CalculateMinutePrices(tradeData)
		model.UpdateCandleData(groupedData)

		// 오래된 데이터 삭제 (1000개 이상 유지하지 않음)
		if len(tradeData) > 1000 {
			tradeData = tradeData[500:]
		}
	}
}
