package controllers

import (
	"fmt"
	"net/http"
	"todo-app/models"
	"todo-app/services"

	"github.com/gin-gonic/gin"
)

func GetTodosByUserId(context *gin.Context) {
	var todos []models.Todo
	userEmail := context.GetHeader("email")

	err := services.GetAllTodosByUserId(&todos, userEmail)

	if err != nil {
		context.AbortWithStatus(http.StatusNotFound)
	} else {
		context.JSON(http.StatusOK, todos)
	}

}

func CreateTodo(context *gin.Context) {
	var todo models.Todo
	email := context.GetHeader("email")

	context.BindJSON(&todo)
	err := services.CreateTodo(&todo, email)
	if err != nil {
		fmt.Println(err.Error())
		context.AbortWithStatus(http.StatusNotFound)
	} else {
		context.JSON(http.StatusOK, todo)
	}
}
