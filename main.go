package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "https://api.bitflyer.jp/v1/ticker?product_code=BTC_JPY"

	req, _ := http.NewRequest(http.MethodGet, url, nil)
	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	} else {
		defer resp.Body.Close()
		v, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		} else {
			var res Responce
			json.Unmarshal(v, &res)
			fmt.Printf("%f", res.Ltp)
		}
	}
}

type Responce struct {
	Type            string  `json:"product_code"`
	State           string  `json:"state"`
	TimeStamp       string  `json:"timestamp"`
	TickID          int     `json:"tick_id"`
	BestBID         float32 `json:"best_bid"`
	BestAsk         float32 `json:"best_ask"`
	BestBIDSize     int     `json:"best_bid_size"`
	BestAskSize     float32 `json:"best_ask_size"`
	TotalBIDDepth   float32 `json:"total_bid_depth"`
	TotalAskDepth   float32 `json:"total_ask_depth"`
	MarketBIDSize   float32 `json:"market_bid_size"`
	MarketAskSize   float32 `json:"market_ask_size"`
	Ltp             float64 `json:"ltp"`
	Volume          float32 `json:"volume"`
	VolumeByProduct float32 `json:"volume_by_product"`
}
