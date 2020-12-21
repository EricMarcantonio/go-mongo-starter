package users

import (
	"context"
)

func AddUser(user User) {
	_, _ = UserCollection.InsertOne(context.TODO(), user)
}

func DeleteUser(user User){
	_, _ = UserCollection.DeleteOne(context.TODO(), user)
}

func GetUser(name string){
	UserCollection.FindOne(context.TODO(), User{Name: name})
}