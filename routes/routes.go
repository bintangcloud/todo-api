package routes

import (
	"todo-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/users", controllers.RegisterUser)

	r.POST("/login", controllers.Login)

	r.GET("/users", controllers.GetAllUsers)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	r.POST("/todos", controllers.CreateTodos)
	r.GET("/todos", controllers.GetAllTodos)
	r.PUT("/todos/:id", controllers.UpdateTodos)
	r.DELETE("/todos/:id", controllers.DeleteTodos)
}
