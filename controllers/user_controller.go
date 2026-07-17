package controllers

import (
	"todo-api/models"

	"github.com/gin-gonic/gin"
)

func CreateUsers(c *gin.Context) {
	r.POST("/users", func(c *gin.Context) {
		var UserBaru models.User
		if err := c.ShouldBindJSON(&UserBaru); err != nil {
			c.JSON(400, gin.H{"error": "Format JSON salah"})
			return
		}

		db.Create(&UserBaru)
		c.JSON(200, gin.H{"status": "User berhasil ditambahkan"})
	})
}

func GetAllUsers(c *gin.Context) {
	r.GET("/users", func(c *gin.Context) {
		var AllUsers []models.User

		db.Find(&AllUsers)
		c.JSON(200, gin.H{"data": AllUsers})
	})
}

func UpdateUser(c *gin.Context) {
	r.PUT("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		var UserLama models.User

		if err := db.First(&UserLama, id).Error; err != nil {
			c.JSON(404, gin.H{"error": "User tidak ditemukan!"})
			return
		}

		var UserBaru models.User
		if err := c.ShouldBindJSON(&UserBaru); err != nil {
			c.JSON(400, gin.H{"error": "Format JSON salah"})
			return
		}

		db.Model(&UserLama).Updates(UserBaru)
		c.JSON(200, gin.H{"status": "User berhasil diupdate!"})
	})
}

func DeleteUser(c *gin.Context) {
	r.DELETE("/users/:id", func(c *gin.Context) {
		id := c.Param("id")

		db.Delete(&models.User{}, id)
		c.JSON(200, gin.H{"status": "sukses menghapus user id" + id})
	})
}
