package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	userBody.ID = uuid.New().String()
	models.Users = append(models.Users, userBody)

	c.JSON(200, gin.H{
		"error": false,
	})

	fmt.Println(models.Users)
}

func EditUser(c *gin.Context) {
	id := c.Param("id")
	var userBody models.User
	err := c.ShouldBindJSON(&userBody)
	if err != nil {
		c.JSON(422, gin.H{
			"error": true,
			"message": "invalid request body",
		})
		return
	}
	for i, u := range models.Users {
		if u.ID == id {
			models.Users[i].Name = userBody.Name
			models.Users[i].Age = userBody.Age
			c.JSON(200, gin.H{
				"error": false,
			})
			return
		}
	}
	c.JSON(404, gin.H{
		"error": true,
		"message": "invalid user id",
	})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	for i, u := range models.Users {
		if u.ID == id {
			models.Users = append(models.Users[:i], models.Users[i + 1:]...)
		}
	}
}

