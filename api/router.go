package api

import (
	"oms-test/internal/user"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	userController := user.UserContorller{}

	router.GET("/user", userController.SearchUser)
	router.POST("/user", userController.CreateUser)
	router.GET("/user/:id", userController.FetchUser)
	router.PUT("/user/:id", userController.UpdateUser)
	router.DELETE("/user/:id", userController.DeleteUser)
}