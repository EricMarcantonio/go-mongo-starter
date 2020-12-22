package server

import (
	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()
	router.POST("/user", HandleAddUser)
	router.GET("/user", HandleHelloWorld)
	router.DELETE("/user")
	router.PUT("/user")
	_ = router.Run()
}
