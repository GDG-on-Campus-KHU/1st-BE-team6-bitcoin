package util

import (
	"log"
	"sort"
	"time"
)

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
func CalculateMinutePrices(data []TradeData) []GroupedData {
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
