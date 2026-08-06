package repositories

import (
	"errors"
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

func FindTodoByID(id string, userID uint) (*models.Todo, error) {

	var todo models.Todo

	err := database.DB.
		Where("id = ? AND user_id = ?", id, userID).
		First(&todo).Error

	return &todo, err
}

func UpdateTodo(todo *models.Todo) error {

	return database.DB.
		Model(todo).
		Updates(map[string]interface{}{
			"title": todo.Title,
		}).Error
}

func DeleteTodo(id string, userID uint) error {

	result := database.DB.
		Where("id = ? AND user_id = ?", id, userID).
		Delete(&models.Todo{})

	if result.RowsAffected == 0 {
		return errors.New("todo tidak ditemukan")
	}

	return result.Error
}
