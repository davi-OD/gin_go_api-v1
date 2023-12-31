package main

import (
	"github.com/davi-OD/go-jwt/controllers"
	"github.com/davi-OD/go-jwt/initializers"
	"github.com/davi-OD/go-jwt/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/signup", controllers.SignUp)

	r.POST("/login", controllers.Login)

	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.Run() // listen and serve on 0.0.0.0:3000
}
