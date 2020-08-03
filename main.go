package main

import (
	"os"
    "os/signal"
	"github.com/gin-gonic/gin"
	"github.com/siesgstarena/epicentre/config"
	"github.com/siesgstarena/epicentre/logger"
	routes "github.com/siesgstarena/epicentre/router"
	"github.com/siesgstarena/epicentre/services/mongo"
	"github.com/siesgstarena/epicentre/services/kafka"
)

func main() {

	err := config.LoadConfig()

	if err != nil {
		panic(err)
	}

	err = logger.LoadLogger(*config.Config)
	if err != nil {
		panic(err)
	}
	logger.Log.Info("Logger Installed Successfully")

	mongo.LoadMongo()

	err = kafka.LoadKafka()
	if err != nil {
		panic(err)
	}
	logger.Log.Info("Kafka Installed Successfully")
	
	go kafka.ConsumeMessage()

	router := gin.Default()

	routes.LoadRouter(router)

	router.Run(":" + config.Config.Port)

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	logger.Log.Info("Shutdown Server ...")
}