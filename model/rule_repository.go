package model

import (
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
		panic(err)
	}

	projectID, err := primitive.ObjectIDFromHex(rule.ProjectID.Hex())
	if err != nil {
		panic(err)
	}

	_, err = mongo.Rules.InsertOne(c, bson.M{
		"userid": userID,
		"projectid": projectID,
		"addon-attachment": rule.HerokuAddonAttachment,
		"addon": rule.HerokuAddon,
		"app": rule.HerokuApp,
		"build": rule.HerokuBuild,
		"collaborator": rule.HerokuCollaborator,
		"domain": rule.HerokuDomain,
		"dyno": rule.HerokuDyno,
		"create": rule.Create,
		"delete": rule.Delete,
		"deployment": rule.Deployment,
		"issues": rule.Issues,
		"issue_comment": rule.IssueComment,
		"pull_request": rule.PR,
		"pull_request_review": rule.PrReview,
		"pull_request_review_comment": rule.PrReviewComment,
		"push": rule.Push,
		"release": rule.Release,
		"project": rule.Project,
		"project_card": rule.ProjectCard,
		"project_column": rule.ProjectColumn,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{"message":"rule Created Sucessfully"})
}

// EditRule Edits rules for project info
func EditRule(c *gin.Context)  {

	var rule Rules 
	c.BindJSON(&rule)

	ruleID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		panic(err)
	}

	filter := bson.M{"_id": ruleID} 

	update := bson.M{
		"$set": bson.M{
			"addon-attachment": rule.HerokuAddonAttachment,
			"addon": rule.HerokuAddon,
			"app": rule.HerokuApp,
			"build": rule.HerokuBuild,
			"collaborator": rule.HerokuCollaborator,
			"domain": rule.HerokuDomain,
			"dyno": rule.HerokuDyno,
			"create": rule.Create,
			"delete": rule.Delete,
			"deployment": rule.Deployment,
			"issues": rule.Issues,
			"issue_comment": rule.IssueComment,
			"pull_request": rule.PR,
			"pull_request_review": rule.PrReview,
			"pull_request_review_comment": rule.PrReviewComment,
			"push": rule.Push,
			"release": rule.Release,
			"project": rule.Project,
			"project_card": rule.ProjectCard,
			"project_column": rule.ProjectColumn,
		},
	}

	result, err := mongo.Rules.UpdateOne(c,filter,update)

	if err != nil {
		panic(err)
	}

	if result.MatchedCount > 0 {
		c.JSON(200, gin.H{"message":"Rule Edited Sucessfully"})
	} else {
		c.JSON(200, gin.H{"message":"No such rule"})
	}
}

// DeleteRule Deletes rule from MongoDB Database
func DeleteRule(c *gin.Context)  {

	ruleID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		panic(err)
	}

	resultRule, err := mongo.Rules.DeleteOne(c, bson.M{"_id": ruleID})
	if err != nil {
		panic(err)
	}

	if resultRule.DeletedCount > 0 {
		c.JSON(200, gin.H{"message":"Rule deleted Sucessfully"})
	} else {
		c.JSON(200, gin.H{"message":"No such rule"})
	}
}

// RuleInfo Gives information of a Rule
func RuleInfo(c *gin.Context)  {

	var rule Rules

	ruleID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		panic(err)
	}

	filter := bson.M{"_id":ruleID}

	if err := mongo.Rules.FindOne(c, filter).Decode(&rule); err != nil {
		panic(err)
	}

	c.JSON(200, rule)
}