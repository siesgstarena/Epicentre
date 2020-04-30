package model 

import "go.mongodb.org/mongo-driver/bson/primitive"

//Rule Type exported for use in API
type Rule struct {
	ID     			primitive.ObjectID 	`bson:"_id,omitempty"`
	User			User 				`bson:"userid,omitempty"`
	Project   		Project 			`bson:"projectid,omitempty"`
	HerokuAddons 	bool				`bson:"herokuaddons,omitempty"`
	HerokuBuilds 	bool				`bson:"herokubuilds,omitempty"`
	CreateBranch 	bool 				`bson:"createbranch,omitempty"`
	DeleteBranch 	bool				`bson:"deletebranch,omitempty"`
	PR 				bool 				`bson:"pr,omitempty"`
	PrReview		bool 				`bson:"prreview,omitempty"`
	Push			bool 				`bson:"push,omitempty"`
	Release			bool 				`bson:"release,omitempty"`
	Issues   		bool				`bson:"issues,omitempty"`
	Forks			bool				`bson:"forks,omitempty"`
	ProjectBoard	bool				`bson:"projectboard,omitempty"`	
	ProjectCard		bool				`bson:"projectcard,omitempty"`	
	ProjectColumn	bool				`bson:"projectcolumn,omitempty"`		
}