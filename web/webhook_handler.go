package web

import (
	"fmt"
	// "net/http"
	"github.com/gin-gonic/gin"
)

// SubscribeWebhook Change Subscription of Webhook
func SubscribeWebhook(c *gin.Context)  {

	var info map[string]interface{}

	c.BindJSON(&info)

	fmt.Println(info["action"])

	c.JSON(204,"Webhook Reveived")
}


// ReceiveWebhooks Accept Data Comming from Heroku as Webhook
func ReceiveWebhooks(c *gin.Context)  {

	var info map[string]interface{}

	c.BindJSON(&info)

	fmt.Println(info)

	c.JSON(204,"Webhook Reveived")
}