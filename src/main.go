package main

import (
	"github.com/gin-gonic/gin"
	"github.com/siesgstarena/epicentre/src/config"
	routes "github.com/siesgstarena/epicentre/src/router"
	"github.com/siesgstarena/epicentre/src/services/logger"
)

func main() {

	router := gin.Default()

	config.LoadConfig(router)

	// access outside logger package
	loggerConfig := logger.Config{
		FileName:   config.Config.FileName,
		MaxSize:    config.Config.MaxSize,
		MaxAge:     config.Config.MaxAge,
		MaxBackUp:  config.Config.MaxBackUp,
		Compress:   config.Config.Compress,
		Level:      config.Config.Level,
		OutputType: config.Config.OutputType,
	}

	err := logger.LoadLogger(loggerConfig)
	if err != nil {
		panic(err)
	}

	logger.Log.Info("Logger Installed Successfully")

	routes.LoadRouter(router)

	router.Run(":" + config.Config.Port)
}
