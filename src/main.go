package main

import (
	"github.com/gin-gonic/gin"
	"github.com/siesgstarena/epicentre/src/config"
)

func main() {

	router := gin.Default()
	
	config.LoadConfig(router)
}
