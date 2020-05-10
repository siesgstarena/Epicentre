package model 

import "go.mongodb.org/mongo-driver/bson/primitive"

type (

	// Project Type exported for use in API
	Project struct {
		ID          	primitive.ObjectID 		`bson:"_id,omitempty"`
		Name       		string             		`bson:"name,omitempty"`
		Description 	string             		`bson:"description,omitempty"`
		Admins    		[]primitive.ObjectID    `bson:"admins,omitempty"`
		Heroku 			Heroku					`bson:"heroku,omitempty"`
		Github			Github 					`bson:"github,omitempty"`
		HealthURL		string					`bson:"healthurl,omitempty"`	
		VersionURL		string					`bson:"versionurl,omitempty"`	
	}

	// Heroku Info regarding project
	Heroku struct {
		AppID			string					`bson:"appID,omitempty"`
		WebhookID 		string					`bson:"webhookID,omitempty"`
	}

	// Github Info regarding project
	Github struct {
		URL				string					`bson:"url,omitempty"`
	}

)