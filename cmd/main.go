package main

import "github.com/gin-gonic/gin"

func getHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Labas",
	})
}

func main() {
	r := gin.Default()

	r.GET("/", getHandler)

	r.Run()
}
