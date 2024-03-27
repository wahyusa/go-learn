package main

import (
	"go-learn/config"
	"go-learn/controllers"

	"github.com/gin-gonic/gin"
)

func init() {
	config.GetEnv()
	config.ConnectionStart()
	config.MigrateDatabase()
}

func main() {
	r := gin.Default()

	// PING
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	users := r.Group("/users")
	{
		users.POST("/register", controllers.StoreUser)
		users.POST("/login", controllers.LoginAttempt)
		// users.PUT("/users", handleUserUpdate)
		// users.DELETE("/users", handleUserDelete)
		// users.GET("/protected", middlewares.IsAuthenticated, controllers.TestProtected)
	}

	// RUN
	r.Run()
}
