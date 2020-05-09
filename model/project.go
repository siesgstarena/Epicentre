package model 

import "go.mongodb.org/mongo-driver/bson/primitive"

//Project Type exported for use in API
type Project struct {
	ID          	primitive.ObjectID 		`bson:"_id,omitempty"`
	Name       		string             		`bson:"name,omitempty"`
	Description 	string             		`bson:"description,omitempty"`
	Admins    		[]primitive.ObjectID    `bson:"admins,omitempty"`
	HerokuAppID 	string					`bson:"herokuappID,omitempty"`
	HerokuWebhookID string					`bson:"herokuwebhookID,omitempty"`
	GithubURL		string					`bson:"githuburl,omitempty"`
	HealthURL		string					`bson:"healthurl,omitempty"`	
	VersionURL		string					`bson:"versionurl,omitempty"`	
}