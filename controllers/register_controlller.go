package controllers

import (
	"todo-api/database"
	"todo-api/models"
	"todo-api/utils"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": "Format JSON salah"})
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": "Gagal mengenkripsi password"})
		return
	}

	user.Password = hashedPassword

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(500, gin.H{
			"error": "Gagal menyimpan data",
		})
		return
	}
	c.JSON(201, gin.H{"status": "Register Berhasil"})
}
