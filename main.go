package main

import (
	"github.com/gin-gonic/gin"
	"github.com/siesgstarena/epicentre/config"
	"github.com/siesgstarena/epicentre/logger"
	routes "github.com/siesgstarena/epicentre/router"
	"github.com/siesgstarena/epicentre/services/kafka"
	"github.com/siesgstarena/epicentre/services/mongo"
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

	kafka.Load()

	mongo.LoadMongo()

	router := gin.Default()

	routes.LoadRouter(router)

	router.Run(":" + config.Config.Port)

	// defer mongo.Client.Disconnect(*mongo.Ctx)
}
