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
	CreateBranch 			bool 				`bson:"createbranch,omitempty"`
	DeleteBranch 			bool				`bson:"deletebranch,omitempty"`
	PR 						bool 				`bson:"pr,omitempty"`
	PrReview				bool 				`bson:"prreview,omitempty"`
	Push					bool 				`bson:"push,omitempty"`
	Release					bool 				`bson:"release,omitempty"`
	Issues   				bool				`bson:"issues,omitempty"`
	Forks					bool				`bson:"forks,omitempty"`
	ProjectBoard			bool				`bson:"projectboard,omitempty"`	
	ProjectCard				bool				`bson:"projectcard,omitempty"`	
	ProjectColumn			bool				`bson:"projectcolumn,omitempty"`		
}