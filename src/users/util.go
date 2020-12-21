package users

import (
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type User struct {
	Name	string	`bson:"name,omitempty"`
}

var (
	UserCollection *mongo.Collection
)

func HelloWorld(resp string){
	log.Println("Hello from go-mongo-starter!")
}

