package model 

import "go.mongodb.org/mongo-driver/bson/primitive"

//Projects Type exported for use in API
type Projects struct {
	ID          primitive.ObjectID 		`bson:"_id,omitempty"`
	Name       	string             		`bson:"name,omitempty"`
	Description string             		`bson:"description,omitempty"`
	Admins    	[]primitive.ObjectID    `bson:"admins,omitempty"`
	GithubURL	string					`bson:"githuburl,omitempty"`	
	HealthURL	string					`bson:"healthurl,omitempty"`	
	VersionURL	string					`bson:"versionurl,omitempty"`	
}