package routes

import (
	"todo-app/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	routes := gin.Default()

	todoRoute := routes.Group("/todos")
	{
		todoRoute.GET("", controllers.GetTodos)
		todoRoute.POST("", controllers.CreateTodo)
	}

	userRoute := routes.Group("/users")
	{
		userRoute.GET("", controllers.GetUsers)
		userRoute.POST("", controllers.CreateUser)
	}
	return routes
}
