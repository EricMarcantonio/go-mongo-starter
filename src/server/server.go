package server

import (
	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()
	router.POST("/user", HandleAddUser)
	router.GET("/user", HandleGetUser)
	router.DELETE("/user", HandleDeleteUser)
	_ = router.Run()

}
