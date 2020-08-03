package model 

import "go.mongodb.org/mongo-driver/bson/primitive"

type (
	//Rule Type exported for use in API
	Rule struct {
		ID     			primitive.ObjectID 	`bson:"_id,omitempty"`
		UserID			primitive.ObjectID 	`bson:"userid,omitempty"`
		ProjectID		primitive.ObjectID	`bson:"projectid,omitempty"`
		Heroku 			HerokuWebhooks		`bson:"heroku,omitempty"`
		Github			GithubWebhooks		`bson:"github,omitempty"`
	}

	// HerokuWebhooks Webhooks from Heroku 
	HerokuWebhooks struct {
		AddonAttachment		bool			`bson:"addonAttachment,omitempty"`
		Addon				bool			`bson:"addon,omitempty"`
		App					bool			`bson:"app,omitempty"`
		Build				bool			`bson:"build,omitempty"`
		Collaborator		bool			`bson:"collaborator,omitempty"`
		Domain				bool			`bson:"domain,omitempty"`
		Dyno				bool			`bson:"dyno,omitempty"`
	}

	// GithubWebhooks Webhooks from Github 
	GithubWebhooks struct {
		CheckRun			bool			`bson:"checkRun,omitempty"`	
		Create				bool 			`bson:"create,omitempty"`
		Delete				bool			`bson:"delete,omitempty"`
		Deployment			bool			`bson:"deployment,omitempty"`
		Issues				bool			`bson:"issues,omitempty"`
		IssueComment		bool			`bson:"issueComment,omitempty"`
		PR					bool 			`bson:"pullRequest,omitempty"`
		PrReview			bool 			`bson:"pullRequestReview,omitempty"`
		PrReviewComment		bool 			`bson:"pullRequestReviewComment,omitempty"`
		Push				bool 			`bson:"push,omitempty"`
		Release				bool 			`bson:"release,omitempty"`
		Project				bool			`bson:"project,omitempty"`	
		ProjectCard			bool			`bson:"projectCard,omitempty"`	
		ProjectColumn		bool			`bson:"projectColumn,omitempty"`		
	}

)