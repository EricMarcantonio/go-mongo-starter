package server

import (
	"github.com/gin-gonic/gin"
	"go-mongo-starter/src/users"
)

// HandleAddUser - Handle the add user request. This is like middleware
func HandleAddUser(ctx *gin.Context) {
	user := users.User{Name: ctx.PostForm("name")}
	if result, err := users.AddUser(user); err != nil {
		ctx.JSON(500, err)
	} else {
		ctx.JSON(200, result.InsertedID)
	}
}

// HandleGetUser - Handle the get user request. Do everything with the request in this function
func HandleGetUser(ctx *gin.Context) {
	user := users.GetUser(ctx.GetString("name"))
	ctx.JSON(200, user)
	//If you need to handle the ctx somewhere else use ctx.Copy()
}

// HandleDeleteUser - Delete the user from Mongo
func HandleDeleteUser(ctx *gin.Context) {
	user := users.User{
		Name: ctx.PostForm("name"),
	}
	//Check you what happens below (go is amazing!) (hint: you don't need a bson, but you should use it)
	if _, err := users.DeleteUser(user); err != nil {
		ctx.JSON(500, err)
	} else {
		ctx.JSON(200, "OK")
	}
}
