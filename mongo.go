package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"
)

// MongoHandler struct a type
type MongoHandler struct {
	client *mongo.Client
	// opts   *options.ClientOptions
}

// Result result struct

// NewMongoHandler creates a new mongohandler
func NewMongoHandler(ip string, port string) MongoHandler {
	opts := options.Client().ApplyURI("mongodb://" + ip + ":" + port)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		fmt.Println(err.Error())
	}
	return MongoHandler{client}
}

// GetUserMon Returns all of a user's data
func (m MongoHandler) GetUserMon(name string) UserData {
	collection := m.client.Database("testing").Collection("nums")
	result := NewUserData()
	filter := bson.M{"name": name}
	err := collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println(err.Error())
	}
	return result
}
