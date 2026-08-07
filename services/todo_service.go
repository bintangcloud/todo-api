package services

import (
	"todo-api/models"
	"todo-api/repositories"
)

func CreateTodo(todo *models.Todo, userID uint) error {

	todo.UserID = userID

	if todo.Status == "" {
		todo.Status = "pending"
	}

	err := repositories.CreateTodo(todo)

	return err
}

func GetTodos(userID uint) ([]models.Todo, error) {

	return repositories.GetTodos(userID)

}

func UpdateTodo(id string, userID uint, title string) error {

	todo, err := repositories.FindTodoByID(id, userID)

	if err != nil {
		return err
	}

	todo.Title = title

	return repositories.UpdateTodo(todo)
}

func DeleteTodo(id string, userID uint) error {

	return repositories.DeleteTodo(id, userID)

}
