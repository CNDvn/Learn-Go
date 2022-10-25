package services

import (
	"todo-app/config"
	"todo-app/models"

	_ "github.com/go-sql-driver/mysql"
)

func GetAllTodos(todos *[]models.Todo) (err error) {
	if err = config.DB.Find(todos).Error; err != nil {
		return err
	}
	return nil
}

func CreateTodo(todo *models.Todo) (err error) {
	if err = config.DB.Create(todo).Error; err != nil {
		return err
	}
	return nil
}
