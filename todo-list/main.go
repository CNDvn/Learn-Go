package main

import (
	"fmt"
	"todo-app/config"
	"todo-app/models"
	"todo-app/routes"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	config.DB, err = gorm.Open("mysql", config.DbURL(config.BuildDBConfig()))
	config.DB.LogMode(true)
	if err != nil {
		fmt.Println("Status:", err)
	}

	defer config.DB.Close()
	config.DB.AutoMigrate(&models.User{})
	config.DB.AutoMigrate(&models.Todo{})

	routes := routes.SetupRouter()

	routes.Run()
}
