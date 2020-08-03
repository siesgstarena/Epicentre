package model

import (
	"github.com/gin-gonic/gin"
	"github.com/siesgstarena/epicentre/services/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateUser Creates & Stores in MongoDB Database
func CreateUser(c *gin.Context)  {

	var user User
	c.BindJSON(&user)

	_, err := mongo.Users.InsertOne(c, bson.M{
		"name": user.Name,
		"email": user.Email,
		"position": user.Position,
	})

	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{"message":"User Created Sucessfully"})
}

// EditUser Edits user profile info
func EditUser(c *gin.Context)  {

	var user User
	c.BindJSON(&user)

	userID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		panic(err)
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
		panic(err)
	}

	if result.MatchedCount > 0 {
		c.JSON(200, gin.H{"message":"User Edited Sucessfully"})
	} else {
		c.JSON(200, gin.H{"message":"No such user"})
	}
}

// DeleteUser Deletes user from MongoDB Database
func DeleteUser(c *gin.Context)  {

	userID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		panic(err)
	}

	resultUser, err := mongo.Users.DeleteOne(c, bson.M{"_id": userID})
	if err != nil {
		panic(err)
	}

	resultRule, err := mongo.Rules.DeleteMany(c, bson.M{"userid": userID})
	if err != nil {
		panic(err)
	}

	filter := bson.M{"admins": bson.M{"$elemMatch": bson.M{"$eq": userID}}}

	resultproject, err := mongo.Projects.UpdateMany(c,filter,bson.M{ "$pull": bson.M{"admins": userID} })
	if err != nil {
		panic(err)
	}

	if resultUser.DeletedCount > 0 || resultRule.DeletedCount > 0 || resultproject.ModifiedCount > 0 {
		c.JSON(200, gin.H{"message":"User deleted Sucessfully"})
	} else {
		c.JSON(200, gin.H{"message":"No such user"})
	}
}

// AllUsers Gives information of a All Users
func AllUsers(c *gin.Context)  {

	filter :=  bson.M{}

	cursor, err := mongo.Users.Find(c, filter)
	if err != nil {
		panic(err)
	}

	var user []bson.M
	if err = cursor.All(c, &user); err != nil {
		panic(err)
	}

	c.JSON(200, user)
}

// UserInfo Gives information of a User
func UserInfo(c *gin.Context)  {

	var user User

	userID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		panic(err)
	}

	filter := bson.M{"_id":userID}

	if err := mongo.Users.FindOne(c, filter).Decode(&user); err != nil {
		panic(err)
	}

	c.JSON(200, user)
}

// RuleForAProjectConnectedToUser Gives information of a Rule for a Project for a perticular User
func RuleForAProjectConnectedToUser(c *gin.Context)  {

	var rule Rule

	userID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		panic(err)
	}

	projectID, err := primitive.ObjectIDFromHex(c.Param("projectid"))
	if err != nil {
		panic(err)
	}

	filter := bson.M{"userid":userID,"projectid":projectID}

	if err := mongo.Rules.FindOne(c, filter).Decode(&rule); err != nil {
		panic(err)
	}

	c.JSON(200, rule)
}