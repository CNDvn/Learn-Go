package controllers

import (
	"fmt"
	"net/http"
	"todo-app/models"
	"todo-app/services"

	"github.com/gin-gonic/gin"
)

func GetTodos(context *gin.Context) {
	var todos []models.Todo

	err := services.GetAllTodos(&todos)

	if err != nil {
		context.AbortWithStatus(http.StatusNotFound)
	} else {
		context.JSON(http.StatusOK, todos)
	}

}

func CreateTodo(context *gin.Context) {
	var todo models.Todo
	context.BindJSON(&todo)
	err := services.CreateTodo(&todo)
	if err != nil {
		fmt.Println(err.Error())
		context.AbortWithStatus(http.StatusNotFound)
	} else {
		context.JSON(http.StatusOK, todo)
	}
}
