package controllers

import (
	"todo-api/database"
	"todo-api/models"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	var UserRegister models.User
	if err := c.ShouldBindJSON(&UserRegister); err != nil {
		c.JSON(400, gin.H{"error": "Format JSON salah"})
		return
	}
	database.DB.Create(&UserRegister)
	c.JSON(200, gin.H{"status": "Register Berhasil"})
}
