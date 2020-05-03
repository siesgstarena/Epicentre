package model

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/siesgstarena/epicentre/services/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateUser Creates & Stores in MongoDB Database
func CreateUser(c *gin.Context)  {

	var user Users
	c.BindJSON(&user)

	_, err := mongo.Users.InsertOne(c, bson.M{
		"name": user.Name,
		"email": user.Email,
		"position": user.Position,
	})

	if err != nil {
		fmt.Println(err)
	}

	c.JSON(200, gin.H{"message":"User Created Sucessfully"})
}

// EditUser Edits user profile info
func EditUser(c *gin.Context)  {

	var user Users
	c.BindJSON(&user)

	userID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	filter := bson.M{"_id": userID} 

	update := bson.M{
		"$set": bson.M{
			"name": user.Name,
			"email": user.Email,
			"position": user.Position,
		},
	}

	result, err := mongo.Users.UpdateOne(c,filter,update)

	if err != nil {
		fmt.Println(err)
	}

	if result.MatchedCount > 0 {
		c.JSON(200, gin.H{"message":"User Edited Sucessfully"})
	} else {
		c.JSON(200, gin.H{"message":"No such user"})
	}
}

// DeleteProject Deletes project from MongoDB Database
func DeleteProject(c *gin.Context)  {

	projectID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	resultRule, err := mongo.Rules.DeleteMany(c, bson.M{"projectid": projectID})
	if err != nil {
		fmt.Println(err)
	}
	
	resultproject, err := mongo.Projects.DeleteOne(c,bson.M{"_id": projectID})
	if err != nil {
		fmt.Println(err)
	}

	if resultRule.DeletedCount > 0 || resultproject.DeletedCount > 0 {
		c.JSON(200, gin.H{"message":"Project deleted Sucessfully"})
	} else {
		c.JSON(200, gin.H{"message":"No such project"})
	}
}

// UserInfo Gives information of a User
func UserInfo(c *gin.Context)  {

	var user Users

	userID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	filter := bson.M{"_id":userID}

	if err := mongo.Users.FindOne(c, filter).Decode(&user); err != nil {
		fmt.Println(err)
	}

	c.JSON(200, user)
}

// RuleForAProjectConnectedToUser Gives information of a Rule for a Project for a perticular User
func RuleForAProjectConnectedToUser(c *gin.Context)  {

	var rule Rules

	userID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	projectID, err := primitive.ObjectIDFromHex(c.Param("projectid"))
	if err != nil {
		fmt.Println(err)
	}

	filter := bson.M{"userid":userID,"projectid":projectID}

	if err := mongo.Rules.FindOne(c, filter).Decode(&rule); err != nil {
		fmt.Println(err)
	}

	c.JSON(200, rule)
}