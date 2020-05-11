package heroku

import (
	"bytes"
	"time"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"github.com/siesgstarena/epicentre/services/mongo"
	"github.com/siesgstarena/epicentre/config"
	"github.com/siesgstarena/epicentre/model"
)

type subscribeResponse struct {
	ID      string		`bson:"id,omitempty"`
	Level   string		`bson:"level,omitempty"`
	URL 	string  	`bson:"url,omitempty"`
}

type appInfo struct {
	ID		string		`json:"id,omitempty"`
	Name	string		`json:"name,omitempty"`
	WebURL	string		`json:"web_url,omitempty"`
}

// GetAllApps Get all info about apps on heroku
func GetAllApps(c *gin.Context)  {

	url := "https://api.heroku.com/apps"
	method := "GET"
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Add("Accept", "application/vnd.heroku+json; version=3.webhooks")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", config.Config.HerokuAPIToken))
	res, err := client.Do(req)
	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	var info []appInfo
	err = json.Unmarshal([]byte(string(body)),&info)
	if err != nil {
		panic(err)
	}

	c.JSON(200,info)
}

// SubscribeWebhook Change Subscription of Webhook
func SubscribeWebhook (c *gin.Context){

	var project model.Project
	projectID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		panic(err)
	}

	if err := mongo.Projects.FindOne(c, bson.M{"_id":projectID}).Decode(&project); err != nil {
		panic(err)
	}

	url := fmt.Sprintf("https://api.heroku.com/apps/%s/webhooks", project.Heroku.AppID )
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
		panic(err)
	}

	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	req.Header.Add("Accept", "application/vnd.heroku+json; version=3.webhooks")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", config.Config.HerokuAPIToken))
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer res.Body.Close()

	var jsonResponse subscribeResponse
	err = json.Unmarshal([]byte(string(body)),&jsonResponse)
	if err != nil {
		fmt.Println(err)
	}
	var webhookID = jsonResponse.ID

	update := bson.M{
		"$set": bson.M{
			"heroku": bson.M {
				"appID": project.Heroku.AppID,
				"webhookID": webhookID,
			},
		},
	}
	result, err := mongo.Projects.UpdateOne(c,bson.M{"_id": projectID},update)
	if result.MatchedCount > 0 {
		c.JSON(200, gin.H{"message":"Webhook Subscribed Sucessfully"})
	} else {
		c.JSON(200, gin.H{"message":"Some error try again"})
	}
}

// DeleteWebhook Delete Webhook for the project
func DeleteWebhook(c *gin.Context) {

	var project model.Project
	projectID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		panic(err)
	}

	if err := mongo.Projects.FindOne(c, bson.M{"_id":projectID}).Decode(&project); err != nil {
		panic(err)
	}

	url := fmt.Sprintf("https://api.heroku.com/apps/%s/webhooks/%s", project.Heroku.AppID , project.Heroku.WebhookID)
	method := "DELETE"
	timeout := time.Duration(5 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	req.Header.Add("Accept", "application/vnd.heroku+json; version=3.webhooks")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", config.Config.HerokuAPIToken))

	res, err := client.Do(req)
	_, err = ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	update := bson.M{
		"$unset": bson.M{
			"heroku.webhookID": "",
		},
	}
	result, err := mongo.Projects.UpdateOne(c,bson.M{"_id": projectID},update)
	if result.MatchedCount > 0 {
		c.JSON(200, gin.H{"message":"Webhook UnSubscribed Sucessfully"})
	} else {
		c.JSON(200, gin.H{"message":"Some error try again"})
	}
}

// ReceiveWebhooks Accept Data Comming from Heroku as Webhook
func ReceiveWebhooks(c *gin.Context)  {

	var info map[string]interface{}

	c.BindJSON(&info)

	fmt.Println(info)

	c.JSON(204,"Webhook Reveived")
}