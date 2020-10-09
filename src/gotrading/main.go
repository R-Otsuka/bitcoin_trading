package main

import (
	"fmt"
	"gotrading/app/models"
	"gotrading/config"
	"gotrading/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	//apiClient := bitflyer.New(config.Config.ApiKey, config.Config.ApiSecret)
	////fmt.Println(apiClient.GetBalane())
	//ticker, _ := apiClient.GetTicker("BTC_USD")
	//fmt.Println(ticker)
	//fmt.Println(ticker.GetMidPrice())
	//fmt.Println(ticker.DateTime())
	//fmt.Println(ticker.TruncateDateTime(time.Second))
	fmt.Println(models.DbConnection)
}
