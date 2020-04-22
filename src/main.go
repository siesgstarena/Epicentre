package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/siesgstarena/epicentre/src/config"
	"github.com/siesgstarena/epicentre/src/services/logger"
	"github.com/siesgstarena/epicentre/src/web"
)

func main() {

	router := gin.Default()

	config.LoadConfig(router)

	fmt.Println(config.Config)

	// access outside logger package
	loggerConfig := logger.Config{
		FileName:config.Config.FileName,
		MaxSize:config.Config.MaxSize,
		MaxAge:config.Config.MaxAge,
		MaxBackUp:config.Config.MaxBackUp,
		Compress:config.Config.Compress,
		Level:config.Config.Level,
		OutputType:config.Config.OutputType,
	}

	fmt.Println(loggerConfig)

	err := logger.NewLogger(loggerConfig)
	if err != nil {
		panic(err)
	}

	logger.Log.Info("Logger Installed Successfully")

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"message": "URL Does not exist",
		})
		logger.Log.Warn("Some one trying URL which does not exist")
	})

	handler := router.Group("/")
	{
		handler.GET("health", web.HeathHandler)
		handler.GET("version", web.VersionHandler)
	}

	router.Run(":" + config.Config.Port)
}
