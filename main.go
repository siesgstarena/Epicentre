package main

import (
	"fmt"
	"os"
    "os/signal"
    "syscall"
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

	sigchan := make(chan os.Signal, 1)
    signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
    go func() {
		<-sigchan
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		kafka.Consumer.Close()
		
        panic("Consumer Received Terminating signal")
	}()
	
	go kafka.ConsumeMessage()

	router := gin.Default()

	routes.LoadRouter(router)

	router.Run(":" + config.Config.Port)

	// defer mongo.Client.Disconnect(*mongo.Ctx)
}