package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/siesgstarena/epicentre/services/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateRule Creates & Stores in MongoDB Database
func CreateRule(c *gin.Context)  {

	var rule Rules 
	c.BindJSON(&rule)

	userID, err := primitive.ObjectIDFromHex(rule.UserID.Hex())
	if err != nil {
		fmt.Println(err)
	}

	projectID, err := primitive.ObjectIDFromHex(rule.ProjectID.Hex())
	if err != nil {
		fmt.Println(err)
	}

	_, err = mongo.Rules.InsertOne(c, bson.D{
		{Key: "userid", Value: userID},
		{Key: "projectid", Value: projectID},
		{Key: "herokuaddons", Value: rule.HerokuAddons},
		{Key: "herokubuilds", Value: rule.HerokuBuilds},
		{Key: "createbranch", Value: rule.CreateBranch},
		{Key: "deletebranch", Value: rule.DeleteBranch},
		{Key: "pr", Value: rule.PR},
		{Key: "prreview", Value: rule.PrReview},
		{Key: "push", Value: rule.Push},
		{Key: "release", Value: rule.Release},
		{Key: "issues", Value: rule.Issues},
		{Key: "forks", Value: rule.Forks},
		{Key: "projectboard", Value: rule.ProjectBoard},
		{Key: "projectcard", Value: rule.ProjectCard},
		{Key: "projectcolumn", Value: rule.ProjectColumn},
	})

	if err != nil {
		fmt.Println(err)
	}

	c.JSON(200, gin.H{"message":"rule Created Sucessfully"})
}
