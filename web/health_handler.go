package web

import (
	"time"
	"github.com/gin-gonic/gin"
	"github.com/siesgstarena/epicentre/services/kafka"
)

// type health struct {
// 	Name     		string `json:"name"`
// 	Description     string `json:"description"`
// 	Health 			string `json:"health"`
// 	Timestamp       time.Time `json:"timestamp"`
// }

// HeathHandler Sends info about health of API
func HeathHandler(c *gin.Context)  {
	info := new(kafka.Health)
	info.Name = "epicentre"
	info.Description = "Cloud Monitoring and Alerting Tool built by SIESGSTarena Platform Team"
	info.Health = "UP"
	loc, _ := time.LoadLocation("UTC")
	info.Timestamp = time.Now().In(loc)
	go kafka.ProduceMessage(*info)
	c.JSON(200, info)
}