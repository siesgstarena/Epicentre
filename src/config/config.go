package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/siesgstarena/epicentre/src/web"
)

import env "github.com/caarlos0/env/v6"

type config struct {
	Port         string           `env:"PORT" envDefault:"8000"`
}

// LoadConfig Loads the config
func LoadConfig(router *gin.Engine)  {
	fmt.Println("Loading Config")

	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"message":"URL Does not exist",
		})
	})

	handler := router.Group("/")
	{
		handler.GET("health", web.HeathHandler)
		handler.GET("version", web.VersionHandler)
	}

	router.Run(":" + cfg.Port)
}