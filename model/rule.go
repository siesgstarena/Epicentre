package model 

import "go.mongodb.org/mongo-driver/bson/primitive"

//Rules Type exported for use in API
type Rules struct {
	ID     					primitive.ObjectID 	`bson:"_id,omitempty"`
	UserID					primitive.ObjectID 	`bson:"userid,omitempty"`
	ProjectID   			primitive.ObjectID	`bson:"projectid,omitempty"`
	HerokuAddonAttachment 	bool				`bson:"addon-attachment,omitempty"`
	HerokuAddon 			bool				`bson:"addon,omitempty"`
	HerokuApp				bool				`bson:"app,omitempty"`
	HerokuBuild 			bool				`bson:"build,omitempty"`
	HerokuCollaborator		bool				`bson:"collaborator,omitempty"`
	HerokuDomain			bool				`bson:"domain,omitempty"`
	HerokuDyno				bool				`bson:"dyno,omitempty"`
	Create					bool 				`bson:"create,omitempty"`
	Delete		 			bool				`bson:"delete,omitempty"`
	Deployment				bool				`bson:"deployment,omitempty"`
	Issues   				bool				`bson:"issues,omitempty"`
	IssueComment   			bool				`bson:"issue_comment,omitempty"`
	PR 						bool 				`bson:"pull_request,omitempty"`
	PrReview				bool 				`bson:"pull_request_review,omitempty"`
	PrReviewComment			bool 				`bson:"pull_request_review_comment,omitempty"`
	Push					bool 				`bson:"push,omitempty"`
	Release					bool 				`bson:"release,omitempty"`
	Project					bool				`bson:"project,omitempty"`	
	ProjectCard				bool				`bson:"project_card,omitempty"`	
	ProjectColumn			bool				`bson:"project_column,omitempty"`		
}