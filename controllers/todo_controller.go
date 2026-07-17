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
