package services

import (
	"todo-app/config"
	"todo-app/models"

	_ "github.com/go-sql-driver/mysql"
)

func GetAllTodosByUserId(todos *[]models.Todo, userEmail string) (err error) {
	var user models.User
	if err := FindUserByEmail(&user, userEmail); err != nil {
		return err
	}

	if err = config.DB.Where("user_id = ?", user.ID).Find(todos).Error; err != nil {
		return err
	}
	return nil
}

func CreateTodo(todo *models.Todo, userEmail string) (err error) {
	var user models.User
	if err := FindUserByEmail(&user, userEmail); err != nil {
		return err
	}
	todo.UserID = user.ID
	if err = config.DB.Create(todo).Error; err != nil {
		return err
	}
	return nil
}
