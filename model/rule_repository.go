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

	_, err = mongo.Rules.InsertOne(c, bson.M{
		"userid": userID,
		"projectid": projectID,
		"herokuaddons": rule.HerokuAddons,
		"herokubuilds": rule.HerokuBuilds,
		"createbranch": rule.CreateBranch,
		"deletebranch": rule.DeleteBranch,
		"pr": rule.PR,
		"prreview": rule.PrReview,
		"push": rule.Push,
		"release": rule.Release,
		"issues": rule.Issues,
		"forks": rule.Forks,
		"projectboard": rule.ProjectBoard,
		"projectcard": rule.ProjectCard,
		"projectcolumn": rule.ProjectColumn,
	})

	if err != nil {
		fmt.Println(err)
	}

	c.JSON(200, gin.H{"message":"rule Created Sucessfully"})
}

// EditRule Edits user profile info
func EditRule(c *gin.Context)  {

	var rule Rules 
	c.BindJSON(&rule)

	ruleID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	userID, err := primitive.ObjectIDFromHex(rule.UserID.Hex())
	if err != nil {
		fmt.Println(err)
	}

	projectID, err := primitive.ObjectIDFromHex(rule.ProjectID.Hex())
	if err != nil {
		fmt.Println(err)
	}

	filter := bson.M{"_id": ruleID} 

	update := bson.M{
		"$set": bson.M{
			"userid": userID,
			"projectid": projectID,
			"herokuaddons": rule.HerokuAddons,
			"herokubuilds": rule.HerokuBuilds,
			"createbranch": rule.CreateBranch,
			"deletebranch": rule.DeleteBranch,
			"pr": rule.PR,
			"prreview": rule.PrReview,
			"push": rule.Push,
			"release": rule.Release,
			"issues": rule.Issues,
			"forks": rule.Forks,
			"projectboard": rule.ProjectBoard,
			"projectcard": rule.ProjectCard,
			"projectcolumn": rule.ProjectColumn,
		},
	}

	result, err := mongo.Rules.UpdateOne(c,filter,update)

	if err != nil {
		fmt.Println(err)
	}

	if result.MatchedCount > 0 {
		c.JSON(200, gin.H{"message":"Rule Edited Sucessfully"})
	} else {
		c.JSON(200, gin.H{"message":"No such rule"})
	}
}