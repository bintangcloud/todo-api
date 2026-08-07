package controllers

import (
	"todo-api/models"
	"todo-api/services"
	"todo-api/utils"

	"github.com/gin-gonic/gin"
)

func CreateTodos(c *gin.Context) {
	userID := c.GetUint("userID")

	var TodoBaru models.Todo

	if err := c.ShouldBindJSON(&TodoBaru); err != nil {
		utils.ErrorResponse(
			c,
			400,
			"validation failed",
			utils.ValidationError(err),
		)
	}

	TodoBaru.UserID = userID

	if err := services.CreateTodo(&TodoBaru, userID); err != nil {
		utils.ErrorResponse(
			c,
			500,
			"Gagal membuat todo",
			nil,
		)
		return
	}

	utils.SuccessResponse(
		c,
		201,
		"Todo berhasil ditambahkan",
		nil,
	)
}

func GetAllTodos(c *gin.Context) {
	userID := c.GetUint("userID")

	var todos []models.Todo

	todos, err := services.GetTodos(userID)
	if err != nil {
		utils.ErrorResponse(
			c,
			500,
			"Gagal mengambil todo",
			nil,
		)
		return
	}

	utils.SuccessResponse(
		c,
		200,
		"Todo berhasil diambil",
		todos,
	)
}

func UpdateTodos(c *gin.Context) {
	userID := c.GetUint("userID")

	id := c.Param("id")

	var body struct {
		Title string `json:"title"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		utils.ErrorResponse(
			c,
			400,
			"validation failed",
			utils.ValidationError(err),
		)
		return
	}

	err := services.UpdateTodo(
		id,
		userID,
		body.Title,
	)

	if err != nil {
		utils.ErrorResponse(
			c,
			404,
			"Todo tidak ditemukan",
			nil,
		)
		return
	}

	utils.SuccessResponse(
		c,
		200,
		"Todo berhasil diupdate",
		nil,
	)
}

func DeleteTodos(c *gin.Context) {
	userID := c.GetUint("userID")

	id := c.Param("id")

	err := services.DeleteTodo(id, userID)

	if err != nil {
		utils.ErrorResponse(
			c,
			404,
			"Todo tidak ditemukan",
			nil,
		)
		return
	}

	utils.SuccessResponse(
		c,
		200,
		"Todo berhasil dihapus",
		nil,
	)
}
