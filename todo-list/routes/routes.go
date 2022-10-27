package routes

import (
	"todo-app/controllers"
	"todo-app/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	routes := gin.Default()

	authRoute := routes.Group("/auth")
	{
		authRoute.POST("/login", controllers.GenerateToken)
	}

	todoRoute := routes.Group("/todos").Use(middlewares.Auth())
	{
		todoRoute.GET("", controllers.GetTodosByUserId)
		todoRoute.POST("", controllers.CreateTodo)
	}

	userRoute := routes.Group("/users")
	{
		userRoute.GET("", controllers.GetUsers)
		userRoute.POST("", controllers.CreateUser)
	}
	return routes
}
