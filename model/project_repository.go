package model

import (
	"github.com/gin-gonic/gin"
	"github.com/siesgstarena/epicentre/services/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	MongoDB "go.mongodb.org/mongo-driver/mongo"
)

// CreateProject Creates & Stores in MongoDB Database
func CreateProject(c *gin.Context)  {

	var project Project
	c.BindJSON(&project)

	_, err := mongo.Projects.InsertOne(c, bson.M{
		"name":project.Name,
		"description":project.Description,
		"admins":project.Admins,
		"herokuappID": project.HerokuAppID,
		"githuburl":project.GithubURL,
		"healthurl":project.HealthURL,
		"versionurl":project.VersionURL,
	})
	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{"message":"Project Created Sucessfully"})
}

// EditProject Edits project details info
func EditProject(c *gin.Context)  {

	var project Project
	c.BindJSON(&project)

	projectID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		panic(err)
	}

	filter := bson.M{"_id": projectID} 

	update := bson.M{
		"$set": bson.M{
			"name":project.Name,
			"description":project.Description,
			"admins":project.Admins,
			"herokuappID": project.HerokuAppID,
			"herokuwebhookID": project.HerokuWebhookID,
			"githuburl":project.GithubURL,
			"healthurl":project.HealthURL,
			"versionurl":project.VersionURL,
		},
	}

	result, err := mongo.Projects.UpdateOne(c,filter,update)

	if err != nil {
		panic(err)
	} else if result.MatchedCount > 0 {
		c.JSON(200, gin.H{"message":"Project Edited Sucessfully"})
	} else {
		c.JSON(200, gin.H{"message":"No such project"})
	}
}

// DeleteProject Deletes project from MongoDB Database
func DeleteProject(c *gin.Context)  {

	projectID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		panic(err)
	}

	resultRule, err := mongo.Rules.DeleteMany(c, bson.M{"projectid": projectID})
	if err != nil {
		panic(err)
	}

	var project Project
	if err := mongo.Projects.FindOne(c, bson.M{"_id":projectID}).Decode(&project); err != nil {
		panic(err)
	}
	
	resultproject, err := mongo.Projects.DeleteOne(c,bson.M{"_id": projectID})
	if err != nil {
		panic(err)
	} else if resultRule.DeletedCount > 0 || resultproject.DeletedCount > 0 {
		c.JSON(200, gin.H{"message":"Project deleted Sucessfully"})
	} else {
		c.JSON(200, gin.H{"message":"No such project"})
	}
}

// ProjectsWhereUserAdmin List All Projects of User in which admin
func ProjectsWhereUserAdmin(c *gin.Context)  {

	userID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		panic(err)
	}

	filter := bson.M{"admins": bson.M{"$elemMatch": bson.M{"$eq": userID}}}

	cursor, err := mongo.Projects.Find(c, filter)
	if err != nil {
		panic(err)
	}

	var allProjects []bson.M

	if err = cursor.All(c, &allProjects); err != nil {
		panic(err)
	}

	c.JSON(200, allProjects)
}

// ProjectInfo Gives information of a Project
func ProjectInfo(c *gin.Context)  {

	var project Project

	projectID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		panic(err)
	}

	filter := bson.M{"_id":projectID}

	if err := mongo.Projects.FindOne(c, filter).Decode(&project); err != nil {
		panic(err)
	}

	c.JSON(200, project)
}

// AllProjects Gives information of a All Projects
func AllProjects(c *gin.Context)  {

	filter :=  bson.M{}

	cursor, err := mongo.Projects.Find(c, filter)
	if err != nil {
		panic(err)
	}

	var project []bson.M
	if err = cursor.All(c, &project); err != nil {
		panic(err)
	}

	c.JSON(200, project)
}

// AllUsersInProject List All Users Monitoring a Project
func AllUsersInProject(c *gin.Context)  {

	projectID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		panic(err)
	}

	matchStage := bson.D{{Key: "$match", Value:bson.D{{Key:"projectid", Value:projectID}}}}
	lookupStage := bson.D{{Key: "$lookup", Value:bson.D{{Key: "from", Value:"users"}, {Key: "localField",Value: "userid"}, {Key: "foreignField",Value: "_id"}, {Key: "as",Value: "userid"}}}}
	unwindStage := bson.D{{Key: "$unwind", Value:bson.D{{Key: "path", Value:"$userid"}}}}

	showLoadedCursor, err := mongo.Rules.Aggregate(c, MongoDB.Pipeline{ matchStage, lookupStage, unwindStage})
	if err != nil {
		panic(err)
	}

	var rules []bson.M
	if err = showLoadedCursor.All(c, &rules); err != nil {
		panic(err)
	}

	c.JSON(200, rules)
}