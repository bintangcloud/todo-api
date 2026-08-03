package routes

import (
	"todo-api/controllers"
	"todo-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/register", controllers.RegisterUser)

	r.POST("/login", controllers.Login)

	r.GET("/users", controllers.GetAllUsers)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	authorized := r.Group("/")
	authorized.Use(middleware.AuthMiddleware())
	authorized.POST("/todos", controllers.CreateTodos)
	authorized.GET("/todos", controllers.GetAllTodos)
	authorized.PUT("/todos/:id", controllers.UpdateTodos)
	authorized.DELETE("/todos/:id", controllers.DeleteTodos)
}
