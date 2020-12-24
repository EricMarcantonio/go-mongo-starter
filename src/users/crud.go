package users

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetUser(name string) *mongo.SingleResult {
	return UserCollection.FindOne(context.TODO(), bson.D{{"name", name}})
}

func AddUser(user User) (*mongo.InsertOneResult, error) {
	return UserCollection.InsertOne(context.TODO(), bson.D{{"name", user.Name}})
}

func DeleteUser(user User) (*mongo.DeleteResult, error) {
	return UserCollection.DeleteMany(context.TODO(), bson.D{{"name", user.Name}})
}
