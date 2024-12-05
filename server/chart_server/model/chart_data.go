package model

import (
	"github.com/Chocobone/GDG_backend_sideProject/server/chart_server/util"
	"sync"
)

type ChartData struct {
	Mu         sync.Mutex
	CandleData [][4]int64 // 캔들 데이터 (시작가, 최고가, 최저가, 종가)
	Timestamps []string   // X축 라벨 (시간)
}

var chartData = &ChartData{}

func UpdateCandleData(groupedData []util.GroupedData) {
	chartData.Mu.Lock()
	defer chartData.Mu.Unlock()

	for _, gd := range groupedData {
		chartData.CandleData = append(chartData.CandleData, [4]int64{
			gd.OpeningPrice, gd.HighPrice, gd.LowPrice, gd.ClosingPrice,
		})
		chartData.Timestamps = append(chartData.Timestamps, gd.Minute)
	}

	if len(chartData.CandleData) > 100 {
		chartData.CandleData = chartData.CandleData[len(chartData.CandleData)-100:]
		chartData.Timestamps = chartData.Timestamps[len(chartData.Timestamps)-100:]
	}
}

func GetChartData() *ChartData {
	return chartData
}
