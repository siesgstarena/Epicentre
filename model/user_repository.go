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

	result, err := mongo.User.InsertOne(*mongo.Ctx, bson.D{
		{Key: "name", Value: "Swapnil Satish Shinde"},
		{Key: "email", Value: "swapnil.satish17@siesgst.ac.in"},
		{Key: "position", Value: "Backend Developer"},
	})

	if err != nil {
		fmt.Println(err)
	}

	c.JSON(200, result.InsertedID)
}
