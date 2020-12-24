package database

/*
	Create your collection connections to different parts, allows for easy manipulation
*/

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	// Client - Global Client ref
	Client *mongo.Client
)

// uses the global client, and creates a collection
func UserCollection() *mongo.Collection {
	return Client.Database("test").Collection("users")
}
