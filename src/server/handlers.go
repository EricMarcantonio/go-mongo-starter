package server

import (
	"github.com/gin-gonic/gin"
	"go-mongo-starter/src/users"
	"go.mongodb.org/mongo-driver/mongo"
)

type FOO struct {
	Name string `json:"url" binding:"required"`
}

func HandleAddUser(ctx *gin.Context) {
	user := users.User{Name: ctx.PostForm("name")}
	if result, err := users.AddUser(user); err != nil {
		ctx.JSON(500, err)
	} else {
		ctx.JSON(200, result.InsertedID)
	}
}

func HandleUpdateUser(ctx *gin.Context) (*mongo.UpdateResult, error) {
	oldUser := users.User{Name: ctx.PostForm("oldName")}
	newUser := users.User{Name: ctx.PostForm("newName")}
	if _, err := users.UpdateUserByName(oldUser, newUser); err != nil {
		ctx.JSON(500, err)
	} else {
		ctx.JSON(200, "OK")
	}
}

func HandleGetUser(ctx *gin.Context) {
	users.GetUser(ctx.GetString("name"))
}

func HandleDeleteUser(ctx *gin.Context) {

}
