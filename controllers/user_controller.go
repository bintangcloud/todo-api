package controllers

import (
	"todo-api/database"
	"todo-api/models"
	"todo-api/utils"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(c *gin.Context) {
	var AllUsers []models.User

	database.DB.Find(&AllUsers)
	utils.SuccessResponse(
		c, 200,
		"Data pengguna",
		gin.H{"data": AllUsers})
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	var UserLama models.User

	if err := database.DB.First(&UserLama, id).Error; err != nil {
		utils.ErrorResponse(c, 404, "User tidak ditemukan!", gin.H{})
		return
	}

	var UserBaru models.User
	if err := c.ShouldBindJSON(&UserBaru); err != nil {
		c.JSON(400, gin.H{"error": "Format JSON salah"})
		return
	}

	database.DB.Model(&UserLama).Updates(UserBaru)
	utils.SuccessResponse(c, 200, "User berhasil diupdate!", gin.H{})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	database.DB.Delete(&models.User{}, id)
	utils.SuccessResponse(c, 200, "sukses menghapus user id"+id, gin.H{})
}
