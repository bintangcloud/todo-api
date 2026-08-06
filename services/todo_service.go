package services

import (
	"todo-api/models"
	"todo-api/repositories"
)

func CreateTodo(todo *models.Todo, userID uint) error {

	todo.UserID = userID

	err := repositories.CreateTodo(todo)

	return err
}
