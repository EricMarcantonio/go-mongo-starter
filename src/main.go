package main

import (
	"context"
	"go-mongo-starter/src/database"
	"go-mongo-starter/src/server"
	"go-mongo-starter/src/users"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func main() {
	//connect to the locally running mongo instance
	client, ctx := connect()
	//close the connection when you are finished with it
	defer client.Disconnect(ctx)

	//Set some globals, cleaner codebase
	database.Client = client
	users.UserCollection = database.UserCollection()
	server.StartServer()
}

func connect() (*mongo.Client, context.Context) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	CheckErr(err)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	CheckErr(err)
	return client, ctx
}
