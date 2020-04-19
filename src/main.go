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

	logger.Load()

	config.LoadConfig(router)

	fmt.Println(config.Config)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"message": "URL Does not exist",
		})
	})

	handler := router.Group("/")
	{
		handler.GET("health", web.HeathHandler)
		handler.GET("version", web.VersionHandler)
	}

	router.Run(":" + config.Config.Port)
}
