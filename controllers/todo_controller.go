package controllers

import (
	"todo-api/database"
	"todo-api/models"

	"github.com/gin-gonic/gin"
)

func CreateTodos(c *gin.Context) {
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
}

func GetAllTodos(c *gin.Context) {
	userID := c.GetUint("userID")

	var todos []models.Todo

	database.DB.Where("user_id = ?", userID).Find(&todos)

	c.JSON(200, gin.H{
		"data": todos,
	})
}

func UpdateTodos(c *gin.Context) {
	userID := c.GetUint("userID")

	id := c.Param("id")

	var TodoLama models.Todo

	if err := database.DB.Where(
		"id = ? AND user_id = ?",
		id,
		userID,
	).First(&TodoLama).Error; err != nil {

		c.JSON(404, gin.H{
			"error": "Todo tidak ditemukan",
		})
		return
	}

	var TodoBaru models.Todo

	if err := c.ShouldBindJSON(&TodoBaru); err != nil {
		c.JSON(400, gin.H{
			"error": "Format JSON salah",
		})
		return
	}

	database.DB.Model(&TodoLama).Updates(map[string]interface{}{
		"title": TodoBaru.Title,
	})

	c.JSON(200, gin.H{
		"status": "Todo berhasil diupdate",
	})
}

func DeleteTodos(c *gin.Context) {
	userID := c.GetUint("userID")
	id := c.Param("id")

	database.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Todo{})
	c.JSON(200, gin.H{"status": "sukses menghapus todo"})
}
