package controllers

import (
	"todo-api/database"
	"todo-api/models"

	"github.com/gin-gonic/gin"
)

func CreateTodos(c *gin.Context) {
	var TodoBaru models.Todo
	if err := c.ShouldBindJSON(&TodoBaru); err != nil {
		c.JSON(400, gin.H{"error": "Format JSON salah"})
		return
	}
	database.DB.Create(&TodoBaru)
	c.JSON(200, gin.H{"status": "Todo berhasil ditambahkan"})
}

func GetAllTodos(c *gin.Context) {
	var AllTodos []models.Todo

	database.DB.Find(&AllTodos)
	c.JSON(200, gin.H{"data": AllTodos})
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
