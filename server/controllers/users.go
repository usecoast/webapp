package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserController(router *gin.RouterGroup) {
	users := router.Group("/users")
	
	users.GET("/", getUser)
}

func getUser(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"user": "example@gmail.com",
	})
}