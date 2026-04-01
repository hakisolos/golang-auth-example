package main

import (
	"example/spark/internals/handlers"
	"example/spark/internals/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
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
	r.GET("/user/delete/:email", handlers.DeleteUser)
	r.POST("/login", handlers.HandleLogin)
	r.GET("/profile", middlewares.AuthMiddleware(), handlers.Me)
	//run server:
	r.Run(":8080")
}
