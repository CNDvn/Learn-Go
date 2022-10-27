package services

import (
	"todo-app/config"
	"todo-app/models"

	_ "github.com/go-sql-driver/mysql"
)

func GetAllUsers(users *[]models.User) (err error) {
	if err = config.DB.Model(&models.User{}).Preload("Todos").Find(&users).Error; err != nil {
		return err
	}
	// if err = config.DB.Find(users).Error; err != nil {
	// 	return err
	// }
	return nil
}

func CreateUser(user *models.User) (err error) {
	if err := user.HashPassword(); err != nil {
		return err
	}
	if err := config.DB.Create(user).Error; err != nil {
		return err
	}

	return nil
}

func FindUserByEmail(user *models.User, email string) (err error) {
	if err = config.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return err
	}
	return nil
}
