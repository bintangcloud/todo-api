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
		utils.ErrorResponse(
			c, 400,
			"Format JSON salah",
			gin.H{})
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		utils.ErrorResponse(
			c, 500,
			"Gagal mengenkripsi password",
			gin.H{})
		return
	}

	user.Password = hashedPassword

	if err := database.DB.Create(&user).Error; err != nil {
		utils.ErrorResponse(
			c, 500,
			"Gagal menyimpan data",
			gin.H{})
		return
	}
	utils.SuccessResponse(
		c, 201,
		"Register Berhasil",
		gin.H{})
}
