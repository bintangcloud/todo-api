package routes

import (
	"todo-api/controllers"
	"todo-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/register", controllers.RegisterUser)

	r.POST("/login", controllers.Login)

	authorized := r.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	authorized.GET("/users", controllers.GetAllUsers)
	authorized.PUT("/users/:id", controllers.UpdateUser)
	authorized.DELETE("/users/:id", controllers.DeleteUser)

	authorized.POST("/todos", controllers.CreateTodos)
	authorized.GET("/todos", controllers.GetAllTodos)
	authorized.PUT("/todos/:id", controllers.UpdateTodos)
	authorized.DELETE("/todos/:id", controllers.DeleteTodos)
}
