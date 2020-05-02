package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Client For MongoDB
var Client *mongo.Client

// Users Collection exported for use in APIs
var Users *mongo.Collection

// Projects Collection exported for use in APIs
var Projects *mongo.Collection

// Rules Collection exported for use in APIs
var Rules *mongo.Collection

// LoadMongo Helps configure mongo
func LoadMongo() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://Swapnil:vijaya26@cluster0-oslju.mongodb.net/test?retryWrites=true&w=majority"))
	if err != nil {
		fmt.Println(err)
	}
	Client = client

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("MongoDB Connected Successfully")

	epicentreDatabase := client.Database("epicentre")
	Users = epicentreDatabase.Collection("users")
	Projects = epicentreDatabase.Collection("projects")
	Rules = epicentreDatabase.Collection("rules")

}
