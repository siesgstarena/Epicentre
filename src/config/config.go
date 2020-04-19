package config

import (
	"github.com/gin-gonic/gin"
	"github.com/siesgstarena/epicentre/src/services/logger"
	env "github.com/caarlos0/env/v6"
)

// Config Available everywhere
var Config *config

type config struct {
	Port         string           `env:"PORT" envDefault:"8000"`
}

// LoadConfig Loads the config
func LoadConfig(router *gin.Engine)  {
	logger.Info("Loading Config")

	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		// logger.Error("%+v\n", err)
	}

	Config = &cfg
}