package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-mongo-starter/src/users"
)

func HandleAddUser(ctx *gin.Context){
	ctx.Params.
	fmt.Println(ctx.GetRawData())
}

func HandleHelloWorld(ctx *gin.Context){
	users.HelloWorld()
}
