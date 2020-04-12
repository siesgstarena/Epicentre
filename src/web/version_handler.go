package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// VersionHandler Sends info about version of API
func VersionHandler(c *gin.Context)  {
	fmt.Println("Inside Health Handler")
	c.JSON(200, gin.H{
		"version": "v2.0",
	})
}