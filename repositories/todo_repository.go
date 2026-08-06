package repositories

import (
	"todo-api/database"
	"todo-api/models"
)

func CreateTodo(todo *models.Todo) error {

	err := database.DB.Create(todo).Error

	return err
}