package users

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetUser(name string) *mongo.SingleResult {
	return UserCollection.FindOne(context.TODO(), User{Name: name})
}

func AddUser(user User) (*mongo.InsertOneResult, error) {
	return UserCollection.InsertOne(context.TODO(), bson.D{{"name", user.Name}})
}

func UpdateUserByName(oldUser User, newUser User) (*mongo.UpdateResult, error) {
	opts := options.Update().SetUpsert(true)
	filter := bson.D{{"name", oldUser.Name}}
	update := bson.D{{"$set", bson.D{{"name", newUser.Name}}}}
	return UserCollection.UpdateOne(context.TODO(), filter, update, opts)
}

func DeleteUser(user User) (*mongo.DeleteResult, error) {
	return UserCollection.DeleteOne(context.TODO(), bson.D{{"name", user.Name}})
}
