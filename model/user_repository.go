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