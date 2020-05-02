package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/siesgstarena/epicentre/services/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateProject Creates & Stores in MongoDB Database
func CreateProject(c *gin.Context)  {

	var project Projects
	c.BindJSON(&project)

	_, err := mongo.Projects.InsertOne(c, bson.M{
		"name":project.Name,
		"description":project.Description,
		"admins":project.Admins,
		"githuburl":project.GithubURL,
		"healthurl":project.HealthURL,
		"versionurl":project.VersionURL,
	})

	if err != nil {
		fmt.Println(err)
	}

	c.JSON(200, gin.H{"message":"Project Created Sucessfully"})
}

// EditProject Edits user profile info
func EditProject(c *gin.Context)  {

	var project Projects
	c.BindJSON(&project)

	projectID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	filter := bson.M{"_id": projectID} 

	update := bson.M{
		"$set": bson.M{
			"name":project.Name,
			"description":project.Description,
			"admins":project.Admins,
			"githuburl":project.GithubURL,
			"healthurl":project.HealthURL,
			"versionurl":project.VersionURL,
		},
	}

	result, err := mongo.Projects.UpdateOne(c,filter,update)

	if err != nil {
		fmt.Println(err)
	}

	if result.MatchedCount > 0 {
		c.JSON(200, gin.H{"message":"Project Edited Sucessfully"})
	} else {
		c.JSON(200, gin.H{"message":"No such project"})
	}
}