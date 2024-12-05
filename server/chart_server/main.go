package main

import (
	"github.com/Chocobone/GDG_backend_sideProject/server/chart_server/chart"
	"github.com/Chocobone/GDG_backend_sideProject/server/chart_server/handler"
	"log"
	"net/http"
)

func main() {
	// 데이터 수신 고루틴 실행
	go chart.FetchBitcoinData()

	// HTTP 서버 설정
	http.HandleFunc("/", handler.ChartHandler)
	log.Println("chart_server가 8080 포트에서 실행 중...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
