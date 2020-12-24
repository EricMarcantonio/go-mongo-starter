package server

import (
	"github.com/gin-gonic/gin"
)

// StartServer - Hook up endpoints, and start the server on port 8080
func StartServer() {
	router := gin.Default()
	router.POST("/user", HandleAddUser)
	router.GET("/user", HandleGetUser)
	router.DELETE("/user", HandleDeleteUser)
	_ = router.Run()

}
