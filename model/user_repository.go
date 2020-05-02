package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/siesgstarena/epicentre/services/mongo"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateUser Creates & Stores in MongoDB Database
func CreateUser(c *gin.Context)  {

	var user Users
	c.BindJSON(&user)

	_, err := mongo.Users.InsertOne(c, bson.D{
		{Key: "name", Value: user.Name},
		{Key: "email", Value: user.Email},
		{Key: "position", Value: user.Position},
	})

	if err != nil {
		fmt.Println(err)
	}

	c.JSON(200, gin.H{"message":"User Created Sucessfully"})
}
