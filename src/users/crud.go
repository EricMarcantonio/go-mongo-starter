package users

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetUser - Logic for getting the user. the ctx stays in the handler as per gin req
func GetUser(name string) *mongo.SingleResult {
	return UserCollection.FindOne(context.TODO(), bson.D{{"name", name}})
}

// AddUser - Add a user to the DB
func AddUser(user User) (*mongo.InsertOneResult, error) {
	return UserCollection.InsertOne(context.TODO(), bson.D{{"name", user.Name}})
}

// DeleteUser - Delete a user from the DB
func DeleteUser(user User) (*mongo.DeleteResult, error) {
	return UserCollection.DeleteMany(context.TODO(), bson.D{{"name", user.Name}})
}
