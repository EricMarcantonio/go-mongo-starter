package users

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// You can add other structs here
type User struct {
	Name string `bson:"name,omitempty"`
}

var (
	// UserCollection - Ref for all collections you may have
	UserCollection *mongo.Collection
)
