package main

import (
	"fmt"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/siesgstarena/epicentre/src/config"
	"github.com/siesgstarena/epicentre/src/web"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	router := gin.Default()
	
	config.Loadconfig()

	handler := router.Group("/")
	{
		handler.GET("health", web.HeathHandler)
		handler.GET("version", web.VersionHandler)
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"message":"URL Does not exist",
		})
	})

	router.Run(":" + os.Getenv("PORT"))
	fmt.Println("Application Running on Port: "+os.Getenv("PORT"))
}
