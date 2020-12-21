package users

import (
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type User struct {
	Name string `bson:"name,omitempty"`
}

var (
	UserCollection *mongo.Collection
)

func HelloWorld() {
	log.Println("Hello from go-mongo-starter!")
}
