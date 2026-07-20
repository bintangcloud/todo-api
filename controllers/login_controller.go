package controllers

import (
	"todo-api/database"
	"todo-api/models"
	"todo-api/utils"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": "Format JSON salah",
		})
		return
	}

	var user models.User

	if err := database.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		if err := utils.CheckPassword(user.Password, req.Password); err != nil {
			c.JSON(401, gin.H{
				"error": "Email atau password salah",
			})
			return
		}
	}
}
