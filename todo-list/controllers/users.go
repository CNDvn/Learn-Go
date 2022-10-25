package controllers

import (
	"fmt"
	"net/http"
	"todo-app/models"
	"todo-app/services"

	"github.com/gin-gonic/gin"
)

func GetUsers(context *gin.Context) {
	var users []models.User

	err := services.GetAllUsers(&users)

	if err != nil {
		context.AbortWithStatus(http.StatusNotFound)
	} else {
		context.JSON(http.StatusOK, users)
	}
}

func CreateUser(context *gin.Context){
	var user models.User
	context.BindJSON(&user)
	err := services.CreateUser(&user)
	if err != nil {
		fmt.Println(err.Error())
		context.AbortWithStatus(http.StatusNotFound)
	} else {
		context.JSON(http.StatusOK, user)
	}
}
