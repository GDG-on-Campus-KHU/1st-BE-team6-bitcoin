package main

import (
	"context"
	"fmt"
	"log"

	// "time"

	proto "github.com/Chocobone/GDG_backend_sideProject/server/pb" // 올바른 경로로 수정
	"google.golang.org/grpc"
)

func main() {
	// gRPC 서버 주소와 포트 설정
	serverAddr := "localhost:50051"

	// gRPC 서버에 연결
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	defer conn.Close()

	// Monitoring_ServiceClient 생성
	client := proto.NewMonitoring_ServiceClient(conn) // NewMonitoring_ServiceClient() 메서드를 사용

	// 요청 메시지 생성
	request := &proto.MonitoringRequest{
		StartMonitoring: "KRW-BTC", // 예시 마켓
	}

	// 서버로 요청 보내기 (스트리밍을 처리)
	stream, err := client.SayHello(context.Background(), request)
	if err != nil {
		log.Fatalf("failed to call SayHello: %v", err)
	}

	// 응답 처리
	for {
		// 스트리밍 응답을 받음
		response, err := stream.Recv()
		if err != nil {
			log.Fatalf("failed to receive response: %v", err)
		}

		// 응답 출력
		fmt.Printf("Market: %s\n", response.GetMarket())
		fmt.Printf("Trade Price: %d\n", response.GetTradePrice())
		fmt.Printf("Trade Time: %s\n", response.GetTradeTime())
		// 추가적으로 필요한 응답 항목을 출력
	}

	// 5초 대기 후 종료 (예시로 요청을 보낸 뒤 잠시 대기)
}
