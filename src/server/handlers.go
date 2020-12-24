package server

import (
	"github.com/gin-gonic/gin"
	"go-mongo-starter/src/users"
)

func HandleAddUser(ctx *gin.Context) {
	user := users.User{Name: ctx.PostForm("name")}
	if result, err := users.AddUser(user); err != nil {
		ctx.JSON(500, err)
	} else {
		ctx.JSON(200, result.InsertedID)
	}
}

func HandleGetUser(ctx *gin.Context) {
	user := users.GetUser(ctx.GetString("name"))
	ctx.JSON(200, user)
}

func HandleDeleteUser(ctx *gin.Context) {
	user := users.User{
		Name: ctx.PostForm("name"),
	}
	if _, err := users.DeleteUser(user); err != nil {
		ctx.JSON(500, err)
	} else {
		ctx.JSON(200, "OK")
	}
}
