package database

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	Client *mongo.Client
)

func UserCollection() *mongo.Collection {
	return Client.Database("test").Collection("users")
}
