package data

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/classAndrew/ApertureServer/pkg/server"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"
)

// MongoHandler struct a type for mongohandler
type MongoHandler struct {
	client            *mongo.Client
	userCollection    *mongo.Collection
	starsysCollection *mongo.Collection
	// opts   *options.ClientOptions
}

// NewMongoHandler creates a new mongohandler
func NewMongoHandler(ip string, port string) *MongoHandler {
	opts := options.Client().ApplyURI("mongodb://" + ip + ":" + port)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, opts)
	userCollection := client.Database("ApertureDB").Collection("players")
	starsysCollection := client.Database("ApertureDB").Collection("starsystems")
	if err != nil {
		fmt.Println(err.Error())
	}
	return &MongoHandler{client, userCollection, starsysCollection}
}

// GetUserMon Returns all of a user's data
func (m *MongoHandler) GetUserMon(name string) server.UserData {
	result := NewUserData()
	filter := bson.M{"name": name}
	err := m.userCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println(err.Error())
		// This will return an empty UserData struct. Check for found by looking for "" name
	}
	return result
}

// InsertUserMon Inserts a newly created user into the players collection
func (m *MongoHandler) InsertUserMon(user *server.UserData) {
	var inface bson.M
	temp, _ := json.Marshal(*user)
	json.Unmarshal(temp, &inface)
	m.userCollection.InsertOne(context.TODO(), inface)
}

// InsertStarSystemMon Inserts a newly created star system into the starsystems collection
func (m *MongoHandler) InsertStarSystemMon(starsys *server.StarSystem) {
	var inface bson.M
	temp, _ := json.Marshal(*starsys)
	json.Unmarshal(temp, &inface)
	m.starsysCollection.InsertOne(context.TODO(), inface)
}

// GetStarSystemMon Returns all of a user's data
func (m *MongoHandler) GetStarSystemMon(name string) *server.StarSystem {
	result := server.NewStarSystem()
	filter := bson.M{"name": name}
	err := m.userCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		fmt.Println(err.Error())
		// This will return an empty UserData struct. Check for found by looking for "" name
	}
	return result
}
