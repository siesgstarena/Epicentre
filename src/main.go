package main

import (
	"github.com/gin-gonic/gin"
	"github.com/siesgstarena/epicentre/src/config"
	routes "github.com/siesgstarena/epicentre/src/router"
	"github.com/siesgstarena/epicentre/src/services/logger"
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

	router := gin.Default()

	routes.LoadRouter(router)

	router.Run(":" + config.Config.Port)
}
