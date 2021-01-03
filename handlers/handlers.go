package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-rest/models"
)

func GetUsers(c *gin.Context) {
	c.JSON(200, models.Users)
}

func CreateUser(c *gin.Context) {
	var userBody models.User
	err := c.ShouldBindJSON(&userBody)
	if err != nil {
		c.JSON(422, gin.H{
			"error": true,
			"message": "invalid request body",
		})
		return
	}
	models.Users = append(models.Users, userBody)

	c.JSON(200, gin.H{
		"error": false,
	})

	fmt.Println(models.Users)
}

