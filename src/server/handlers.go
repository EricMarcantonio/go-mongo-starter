package server

import (
	"github.com/gin-gonic/gin"
	"go-mongo-starter/src/users"
)

type FOO struct {
	Name string `json:"url" binding:"required"`
}

func HandleAddUser(ctx *gin.Context) {
	ctx.JSON(200, ctx.PostForm("test"))
}

func HandleHelloWorld(ctx *gin.Context) {
	users.HelloWorld()

}
