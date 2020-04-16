package web

import (
	"fmt"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

type version struct {
    Name     		string `json:"name"`
	Description     string `json:"description"`
	Version 		string `json:"health"`
	Timestamp       time.Time `json:"timestamp"`
}

// VersionHandler Sends info about version of API
func VersionHandler(c *gin.Context)  {
	fmt.Println("Inside Health Handler")
	info := new(version)
	info.Name = "epicentre"
	info.Description = "Cloud Monitoring and Alerting Tool built by SIESGSTarena Platform Team)"
	info.Version = runtime.Version()
	loc, _ := time.LoadLocation("UTC")
    info.Timestamp = time.Now().In(loc)
	c.JSON(200, info)
}