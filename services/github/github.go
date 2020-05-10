package github

import (
	// "bytes"
	// "time"
	// "encoding/json"
	"fmt"
	// "io/ioutil"
	// "net/http"
	"github.com/gin-gonic/gin"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "github.com/siesgstarena/epicentre/services/mongo"
	// "github.com/siesgstarena/epicentre/config"
	// "github.com/siesgstarena/epicentre/model"
)

// ReceiveWebhooks Accept Data Comming from Github as Webhook
func ReceiveWebhooks(c *gin.Context)  {

	var info map[string]interface{}

	c.BindJSON(&info)

	fmt.Println(c.Request.Header.Get("X-GitHub-Event"))

	fmt.Println(info)

	c.JSON(204,"Webhook Reveived")
}