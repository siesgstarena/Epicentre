package github

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
	ID      int		`bson:"id,omitempty"`
	Name   	string	`bson:"name,omitempty"`
	Active 	bool	`bson:"active,omitempty"`
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

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/hooks",project.Github.Owner,project.Github.RepoName)
	method := "POST"

	message := map[string]interface{}{
		"name": "web",
		"config": map[string]interface{} {
			"url": fmt.Sprintf("%s/webhook/github", config.Config.DeployedAppURL),
			"content_type":"json",
		},
		"events": []string {
			"check_run","create","delete","deployment","issues","issue_comment","project_card","project_column","project","pull_request","pull_request_review","pull_request_review_comment","push","release",
		},
		"active": true,
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
	req.Header.Add("Accept", "application/vnd.github.machine-man-preview+json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", config.Config.GithubAPIToken))
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
	fmt.Println(webhookID)

	update := bson.M{
		"$set": bson.M{
			"github": bson.M {
				"owner": project.Github.Owner,
				"repoName": project.Github.RepoName,
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

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/hooks/%d",project.Github.Owner,project.Github.RepoName,project.Github.WebhookID)
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
	req.Header.Add("Accept", "application/vnd.github.machine-man-preview+json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", config.Config.GithubAPIToken))

	res, err := client.Do(req)
	_, err = ioutil.ReadAll(res.Body)
	defer res.Body.Close()

	update := bson.M{
		"$unset": bson.M{
			"github.webhookID": "",
		},
	}
	result, err := mongo.Projects.UpdateOne(c,bson.M{"_id": projectID},update)
	if result.MatchedCount > 0 {
		c.JSON(200, gin.H{"message":"Webhook UnSubscribed Sucessfully"})
	} else {
		c.JSON(200, gin.H{"message":"Some error try again"})
	}
}

// ReceiveWebhooks Accept Data Comming from Github as Webhook
func ReceiveWebhooks(c *gin.Context)  {

	var info map[string]interface{}

	c.BindJSON(&info)

	fmt.Println(c.Request.Header.Get("X-GitHub-Event"))

	fmt.Println(info)

	c.JSON(204,"Webhook Reveived")
}