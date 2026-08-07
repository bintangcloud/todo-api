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
		utils.ErrorResponse(
			c, 400,
			"Format JSON salah",
			gin.H{})
		return
	}

	var user models.User

	if err := database.DB.Where("email = ?", req.Email).First(&user).Error; err != nil {
		if err := utils.CheckPassword(user.Password, req.Password); err != nil {
			utils.ErrorResponse(
				c, 401,
				"Email atau password salah",
				gin.H{})
			return
		}
	}

	token, err := utils.GenerateToken(user)
	if err != nil {
		utils.ErrorResponse(
			c, 500,
			"Gagal membuat token",
			gin.H{})
		return
	}

	utils.SuccessResponse(
		c, 200,
		"Login berhasil",
		gin.H{
			"token": token,
		})

}
