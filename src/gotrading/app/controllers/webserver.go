package controllers

import (
	"fmt"
	"gotrading/app/models"
	"html/template"
	"net/http"

	"gotrading/config"
)

var templates = template.Must(template.ParseFiles("app/views/google.html"))

func viewChartHandler(w http.ResponseWriter, r *http.Request){
	limit := 100
	duration := "1m"
	durationTime := config.Config.Durations[duration]
	df, errr := models.GetAllCandle(config.Config.ProductCode, durationTime, limit)
	if errr != nil{
		http.Error(w, errr.Error(), http.StatusInternalServerError)
	}
	err := templates.ExecuteTemplate(w, "google.html", df.Candles)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func StartWebServer() error {
	http.HandleFunc("/chart/", viewChartHandler)
	return http.ListenAndServe(fmt.Sprintf(":%d",config.Config.Port), nil)
}