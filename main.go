package main

import (
	"log"

	"github.com/kamio06/gotrading/app/controllers"
	"github.com/kamio06/gotrading/config"
	"github.com/kamio06/gotrading/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	controllers.StreamIngestionData()
	log.Println(controllers.StartWebServer())
}
