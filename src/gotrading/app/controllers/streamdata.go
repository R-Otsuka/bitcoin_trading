package controllers

import (
	"gotrading/app/models"
	"gotrading/bitflyer"
	"gotrading/config"
	"log"

	"fmt"
)

func StreamIngestionData() {
	limit := 100
	duration := "1m"
	durationTime := config.Config.Durations[duration]
	df, _ := models.GetAllCandle(config.Config.ProductCode, durationTime, limit)
	log.Printf("df=%v",df)
	for i:=0;i<len(df.Candles);i++{
		//log.Printf("dfCandle=%v\n",df.Candles[i])
		log.Printf("%+v\n", df.Candles[i])
	}
	var tickerChannl = make(chan bitflyer.Ticker)
	apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	go apiClient.GetRealTimeTicker(config.Config.ProductCode, tickerChannl)
	go func(){
		for ticker := range tickerChannl {
			log.Printf("action=StreamIngestionData, %v", ticker)
			for _, duration := range config.Config.Durations {
				isCreated := models.CreateCandleWithDuration(ticker, ticker.ProductCode, duration)
				fmt.Println(isCreated)
				if isCreated == true && duration == config.Config.TradeDuration {
					// TODO
				}
			}
		}
	}()
}
