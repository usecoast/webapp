package main

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"

	controllers "server/controllers"
)

var router = gin.Default()

func main() {
	v1 := router.Group("/api/v1")
	controllers.UserController(v1)
	
	router.Use(static.Serve("/", static.LocalFile("../client/build", true)))
	router.NoRoute(func(c *gin.Context) {
		if !strings.HasPrefix(c.Request.RequestURI, "/api/v1") {
			fmt.Println("Not a valid api route, redirect to client side")
			c.File("../client/build/index.html")
		}
	})

	router.Run(":4000")
}