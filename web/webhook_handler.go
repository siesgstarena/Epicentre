package web

import (
	"bytes"
	"time"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/siesgstarena/epicentre/config"
)

// SubscribeHerokuWebhook Change Subscription of Webhook
func SubscribeHerokuWebhook(AppID string) error  {

	url := fmt.Sprintf("https://api.heroku.com/apps/%s/webhooks", AppID)
	method := "POST"
	message := map[string]interface{}{
		"include": []string {
		  "api:addon-attachment","api:addon","api:build","api:collaborator","api:domain","api:dyno",
		},
		"level": "sync",
		"url": fmt.Sprintf("%s/webhook/heroku", config.Config.DeployedAppURL),
	}
	payload, err := json.Marshal(message)
	if err != nil {
		fmt.Println(err)
		return err
	}

	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Add("Accept", "application/vnd.heroku+json; version=3.webhooks")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", config.Config.HerokuAPIToken))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	fmt.Println(string(body))
	return nil
}


// ReceiveWebhooks Accept Data Comming from Heroku as Webhook
func ReceiveWebhooks(c *gin.Context)  {

	var info map[string]interface{}

	c.BindJSON(&info)

	fmt.Println(info)

	c.JSON(204,"Webhook Reveived")
}