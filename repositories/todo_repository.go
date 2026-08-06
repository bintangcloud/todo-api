package repositories

import (
	"todo-api/database"
	"todo-api/models"
)

func CreateTodo(todo *models.Todo) error {

	err := database.DB.Create(todo).Error

	return err
}

func GetTodos(userID uint) ([]models.Todo, error) {

	var todos []models.Todo

	err := database.DB.
		Where("user_id = ?", userID).
		Find(&todos).Error

	return todos, err
}
