package main

import (
	"github.com/gin-gonic/gin"
	"go-gin-rest/handlers"
	"log"
)

func main() {
	r := gin.Default()

	userRoutes := r.Group("/users")
	{
		userRoutes.GET("/", handlers.GetUsers)
		userRoutes.POST("/", handlers.CreateUser)
		userRoutes.PUT("/:id", handlers.EditUser)
	}

	if err := r.Run(); err != nil {
		log.Fatal(err)
	}
}
