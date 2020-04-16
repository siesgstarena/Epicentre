package main

import (
	"github.com/gin-gonic/gin"
	"github.com/siesgstarena/epicentre/src/config"
	"github.com/siesgstarena/epicentre/src/services/logger"
)

func main() {

	router := gin.Default()

	logger.Load()
	
	config.LoadConfig(router)
}
