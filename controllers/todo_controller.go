package controllers

import (
	"todo-api/database"
	"todo-api/models"

	"github.com/gin-gonic/gin"
)

userID := c.GetUint("userID")

	var TodoBaru models.Todo

	if err := c.ShouldBindJSON(&TodoBaru); err != nil {
		c.JSON(400, gin.H{
			"error": "Format JSON salah",
		})
		return
	}

	TodoBaru.UserID = userID

	if err := database.DB.Create(&TodoBaru).Error; err != nil {
		c.JSON(500, gin.H{
			"error": "Gagal menyimpan todo",
		})
		return
	}

	c.JSON(201, gin.H{
		"status": "Todo berhasil ditambahkan",
	})

func GetAllTodos(c *gin.Context) {
	userID := c.GetUint("userID")

	var todos []models.Todo

	database.DB.Where("user_id = ?", userID).Find(&todos)

	c.JSON(200, gin.H{
		"data": todos,
	})
}

func UpdateTodos(c *gin.Context) {
	id := c.Param("id")
	var TodoLama models.Todo

	if err := database.DB.First(&TodoLama, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Todo tidak ditemukan!"})
		return
	}

	var TodoBaru models.Todo
	if err := c.ShouldBindJSON(&TodoBaru); err != nil {
		c.JSON(400, gin.H{"error": "Format JSON salah"})
		return
	}

	database.DB.Model(&TodoLama).Updates(TodoBaru)
	c.JSON(200, gin.H{"status": "Todo berhasil diupdate!"})
}

func DeleteTodos(c *gin.Context) {
	id := c.Param("id")

	database.DB.Delete(&models.Todo{}, id)
	c.JSON(200, gin.H{"status": "sukses menghapus todo id " + id})
}
