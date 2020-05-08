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

type subscribeResponse struct {
	ID      string		`bson:"id,omitempty"`
	Level   string		`bson:"level,omitempty"`
	URL 	string  	`bson:"url,omitempty"`
}

// SubscribeHerokuWebhook Change Subscription of Webhook
func SubscribeHerokuWebhook(AppID string) ( string, error)  {

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
		return "", err
	}

	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	req.Header.Add("Accept", "application/vnd.heroku+json; version=3.webhooks")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", config.Config.HerokuAPIToken))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()

	fmt.Println(string(body))

	var jsonResponse subscribeResponse

	err = json.Unmarshal([]byte(string(body)),&jsonResponse)
	if err != nil {
		fmt.Println(err)
	}

	var WebhookID = jsonResponse.ID

	return WebhookID,nil
}

// DeleteWebhook Delete Webhook for the project
func DeleteWebhook(AppID string, WebhookID string) error  {

	url := fmt.Sprintf("https://api.heroku.com/apps/%s/webhooks/%s", AppID, WebhookID)
	method := "DELETE"
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Add("Accept", "application/vnd.heroku+json; version=3.webhooks")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", config.Config.HerokuAPIToken))

	res, err := client.Do(req)
	_, err = ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	return nil
}


// ReceiveWebhooks Accept Data Comming from Heroku as Webhook
func ReceiveWebhooks(c *gin.Context)  {

	var info map[string]interface{}

	c.BindJSON(&info)

	fmt.Println(info)

	c.JSON(204,"Webhook Reveived")
}