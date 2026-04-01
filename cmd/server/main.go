package main

import (
	"example/spark/internals/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "sparkdb api running",
		})
	})
	// auth handlers:
	r.POST("/signup", handlers.CreateUser)
	r.GET("/users", handlers.GetUsers)
	r.GET("/user/:id", handlers.GetOneUser)
	//run server:
	r.Run(":8080")
}
