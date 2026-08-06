package controllers

import (
	"todo-api/database"
	"todo-api/models"
	"todo-api/services"

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

	if err := services.CreateTodo(&TodoBaru, userID); err != nil {
		c.JSON(500, gin.H{
			"error": "Gagal membuat todo",
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

	todos, err := services.GetTodos(userID)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Gagal mengambil todo",
		})
		return
	}

	c.JSON(200, gin.H{
		"data": todos,
	})
}

func UpdateTodos(c *gin.Context) {
	userID := c.GetUint("userID")

	id := c.Param("id")

	var body struct {
		Title string `json:"title"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{
			"error": "format salah",
		})
		return
	}

	err := services.UpdateTodo(
		id,
		userID,
		body.Title,
	)

	if err != nil {
		c.JSON(404, gin.H{
			"error": "Todo tidak ditemukan",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": "Todo berhasil diupdate",
	})
}

func DeleteTodos(c *gin.Context) {
	userID := c.GetUint("userID")
	id := c.Param("id")

	result := database.DB.
		Where("id = ? AND user_id = ?", id, userID).
		Delete(&models.Todo{})

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{
			"error": "Todo tidak ditemukan",
		})
		return
	}

	c.JSON(200, gin.H{
		"status": "sukses menghapus todo",
	})
}
