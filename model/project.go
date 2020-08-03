package model 

import "go.mongodb.org/mongo-driver/bson/primitive"

type (

	// Project Type exported for use in API
	Project struct {
		ID          	primitive.ObjectID 		`bson:"_id,omitempty"`
		Name       		string             		`bson:"name,omitempty"`
		Description 	string             		`bson:"description,omitempty"`
		Admins    		[]primitive.ObjectID    `bson:"admins,omitempty"`
		Heroku 			HerokuDetails					`bson:"heroku,omitempty"`
		Github			GithubDetails 					`bson:"github,omitempty"`
		HealthURL		string					`bson:"healthurl,omitempty"`	
		VersionURL		string					`bson:"versionurl,omitempty"`	
	}

	// HerokuDetails Info regarding project
	HerokuDetails struct {
		AppID			string					`bson:"appID,omitempty"`
		WebhookID 		string					`bson:"webhookID,omitempty"`
	}

	// GithubDetails Info regarding project
	GithubDetails struct {
		Owner 			string 					`bson:"owner,omitempty"`
		RepoName 		string 					`bson:"repoName,omitempty"`
		WebhookID 		int					`bson:"webhookID,omitempty"`
	}

)