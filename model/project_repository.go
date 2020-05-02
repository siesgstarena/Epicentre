package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/siesgstarena/epicentre/services/mongo"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateProject Creates & Stores in MongoDB Database
func CreateProject(c *gin.Context)  {

	var project Project
	c.BindJSON(&project)

	_, err := mongo.Project.InsertOne(c, bson.D{
		{Key: "name", Value: project.Name},
		{Key: "description", Value: project.Description},
		{Key: "admins", Value: project.Admins},
		{Key: "githuburl", Value: project.GithubURL},
		{Key: "healthurl", Value: project.HealthURL},
		{Key: "versionurl", Value: project.VersionURL},
	})

	if err != nil {
		fmt.Println(err)
	}

	c.JSON(200, gin.H{"message":"Project Created Sucessfully"})
}
