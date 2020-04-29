package model 

import "go.mongodb.org/mongo-driver/bson/primitive"

//User Type exported for use in API
type User struct {
	ID     		primitive.ObjectID 	`bson:"_id,omitempty"`
	Name  		string             	`bson:"name,omitempty"`
	Email 		string             	`bson:"author,omitempty"`
	Position   	string           	`bson:"tags,omitempty"`
}