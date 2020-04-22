package config

import (
	"fmt"
	"github.com/gin-gonic/gin"
	env "github.com/caarlos0/env/v6"
)

// Config Available everywhere
var Config *config

type (
	config struct {
		Port       string			`env:"PORT" envDefault:"8000"`
		FileName   string			`env:"FILENAME" envDefault:"zap.log"`			
		MaxSize    int				`env:"MAXSIZE" envDefault:"100"`
		MaxAge     int				`env:"MAXAGE" envDefault:"10"`
		MaxBackUp  int				`env:"MAZBACKUP" envDefault:"5"`
		Compress   bool				`env:"COMPRESS" envDefault:"false"`
		Level      string			`env:"LEVEL" envDefault:"info"`
		OutputType string			`env:"OUTPUTTYPE" envDefault:"json"`
	}
)

// LoadConfig Loads the config
func LoadConfig(router *gin.Engine)  {
	fmt.Println("Initializing LoadConfig")

	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		panic(err)
	}

	// Address of parsed configs passing to Global Config variable
	Config = &cfg

	// Configuring Static Assets for development
	router.Static("/assets", "../images")
	router.StaticFile("/favicon.ico", "../images/favicon.ico")

	fmt.Println("Initialization Finished")
}