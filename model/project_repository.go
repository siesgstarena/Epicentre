package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/siesgstarena/epicentre/services/mongo"
	"github.com/siesgstarena/epicentre/web"
	MongoDB "go.mongodb.org/mongo-driver/mongo"
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
		"herokuappID": project.HerokuAppID,
		"githuburl":project.GithubURL,
		"healthurl":project.HealthURL,
		"versionurl":project.VersionURL,
	})
	if err != nil {
		panic(err)
	}

	err = web.SubscribeHerokuWebhook(project.HerokuAppID)
	if err != nil {
		panic(err)
	}
	fmt.Println("Webhook Subscribed Successfully")

	c.JSON(200, gin.H{"message":"Project Created & Subscribed Sucessfully"})
}

// EditProject Edits project details info
func EditProject(c *gin.Context)  {

	var project Projects
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
			"githuburl":project.GithubURL,
			"healthurl":project.HealthURL,
			"versionurl":project.VersionURL,
		},
	}

	result, err := mongo.Projects.UpdateOne(c,filter,update)

	if err != nil {
		panic(err)
	}

	if result.MatchedCount > 0 {
		c.JSON(200, gin.H{"message":"Project Edited Sucessfully"})
	} else {
		c.JSON(200, gin.H{"message":"No such project"})
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

	var project Projects

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