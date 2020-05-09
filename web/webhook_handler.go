package web

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

// SubscribeHerokuWebhook Change Subscription of Webhook
func SubscribeHerokuWebhook (c *gin.Context){

	var project model.Projects
	projectID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		panic(err)
	}

	if err := mongo.Projects.FindOne(c, bson.M{"_id":projectID}).Decode(&project); err != nil {
		panic(err)
	}

	url := fmt.Sprintf("https://api.heroku.com/apps/%s/webhooks", project.HerokuAppID)
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
			"herokuwebhookID": webhookID,
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

	var project model.Projects
	projectID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		panic(err)
	}

	if err := mongo.Projects.FindOne(c, bson.M{"_id":projectID}).Decode(&project); err != nil {
		panic(err)
	}

	url := fmt.Sprintf("https://api.heroku.com/apps/%s/webhooks/%s", project.HerokuAppID, project.HerokuWebhookID)
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
			"herokuwebhookID": "",
		},
	}
	result, err := mongo.Projects.UpdateOne(c,bson.M{"_id": projectID},update)
	if result.MatchedCount > 0 {
		c.JSON(200, gin.H{"message":"Webhook UnSubscribed Sucessfully"})
	} else {
		c.JSON(200, gin.H{"message":"Some error try again"})
	}
}

// ReceiveHerokuWebhooks Accept Data Comming from Heroku as Webhook
func ReceiveHerokuWebhooks(c *gin.Context)  {

	var info map[string]interface{}

	c.BindJSON(&info)

	fmt.Println(info)

	c.JSON(204,"Webhook Reveived")
}

// ReceiveGithubWebhooks Accept Data Comming from Github as Webhook
func ReceiveGithubWebhooks(c *gin.Context)  {

	var info map[string]interface{}

	c.BindJSON(&info)

	fmt.Println(c.Request.Header.Get("X-GitHub-Event"))

	fmt.Println(info)

	c.JSON(204,"Webhook Reveived")
}