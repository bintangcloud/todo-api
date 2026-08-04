package main

import (
	"todo-api/database"
	"todo-api/routes"
	"todo-api/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	utils.LoadJWTSecret()

	database.ConnectDB()

	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":8080")
}
