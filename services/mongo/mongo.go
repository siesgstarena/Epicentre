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

// Ctx For MongoDB 
var Ctx *context.Context

// User Collection exported for use in APIs
var User *mongo.Collection

// Project Collection exported for use in APIs
var Project *mongo.Collection

// Rules Collection exported for use in APIs
var Rules *mongo.Collection

// LoadMongo Helps configure mongo
func LoadMongo()  {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://Swapnil:vijaya26@cluster0-oslju.mongodb.net/test?retryWrites=true&w=majority"))
	if err != nil {
		fmt.Println(err)
	}
	Client = client

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		fmt.Println(err)
	}
	Ctx = &ctx

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("MongoDB Connected Successfully")

	epicentreDatabase := client.Database("epicentre")
    User = epicentreDatabase.Collection("users")
	Project = epicentreDatabase.Collection("projects")
	Rules = epicentreDatabase.Collection("rules")

}