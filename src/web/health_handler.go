package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

// HeathHandler Sends info about health of API
func HeathHandler(c *gin.Context)  {
	fmt.Println("Inside Health Handler")
	c.JSON(200, gin.H{
		"health": "perfect",
	})
}