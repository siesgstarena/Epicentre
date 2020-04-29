package model 

import "go.mongodb.org/mongo-driver/bson/primitive"

//Rule Type exported for use in API
type Rule struct {
	ID     		primitive.ObjectID `bson:"_id,omitempty"`
	UserID		primitive.ObjectID `bson:"userid,omitempty"`
	ProjectID   primitive.ObjectID `bson:"projectid,omitempty"`
	// Rules for notification to be added
}