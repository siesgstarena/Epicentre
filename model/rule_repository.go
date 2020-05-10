package model

import (
	"github.com/gin-gonic/gin"
	"github.com/siesgstarena/epicentre/services/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateRule Creates & Stores in MongoDB Database
func CreateRule(c *gin.Context)  {

	var rule Rule 
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
		"heroku": bson.M {
			"addonAttachment": rule.Heroku.AddonAttachment,
			"addon": rule.Heroku.Addon,
			"app": rule.Heroku.App,
			"build": rule.Heroku.Build,
			"collaborator": rule.Heroku.Collaborator,
			"domain": rule.Heroku.Domain,
			"dyno": rule.Heroku.Dyno,
		},
		"github": bson.M {
			"checkRun": rule.Github.CheckRun,
			"create": rule.Github.Create,
			"delete": rule.Github.Delete,
			"deployment": rule.Github.Deployment,
			"issues": rule.Github.Issues,
			"issueComment": rule.Github.IssueComment,
			"pullRequest": rule.Github.PR,
			"pullRequestReview": rule.Github.PrReview,
			"pullRequestReviewComment": rule.Github.PrReviewComment,
			"push": rule.Github.Push,
			"release": rule.Github.Release,
			"project": rule.Github.Project,
			"projectCard": rule.Github.ProjectCard,
			"projectColumn": rule.Github.ProjectColumn,
		},
	})

	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{"message":"rule Created Sucessfully"})
}

// EditRule Edits rules for project info
func EditRule(c *gin.Context)  {

	var rule Rule 
	c.BindJSON(&rule)

	ruleID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		panic(err)
	}

	filter := bson.M{"_id": ruleID} 

	update := bson.M{
		"$set": bson.M{
			"heroku": bson.M {
				"addonAttachment": rule.Heroku.AddonAttachment,
				"addon": rule.Heroku.Addon,
				"app": rule.Heroku.App,
				"build": rule.Heroku.Build,
				"collaborator": rule.Heroku.Collaborator,
				"domain": rule.Heroku.Domain,
				"dyno": rule.Heroku.Dyno,
			},
			"github": bson.M {
				"checkRun": rule.Github.CheckRun,
				"create": rule.Github.Create,
				"delete": rule.Github.Delete,
				"deployment": rule.Github.Deployment,
				"issues": rule.Github.Issues,
				"issueComment": rule.Github.IssueComment,
				"pullRequest": rule.Github.PR,
				"pullRequestReview": rule.Github.PrReview,
				"pullRequestReviewComment": rule.Github.PrReviewComment,
				"push": rule.Github.Push,
				"release": rule.Github.Release,
				"project": rule.Github.Project,
				"projectCard": rule.Github.ProjectCard,
				"projectColumn": rule.Github.ProjectColumn,
			},
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

	var rule Rule

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