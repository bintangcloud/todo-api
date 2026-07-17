package controllers

import (
	"github.com/gin-gonic/gin"
)

func CreateUsers(c *gin.Context) {
	r.POST("/users", func(c *gin.Context) {
		var UserBaru User
		if err := c.ShouldBindJSON(&UserBaru); err != nil {
			c.JSON(400, gin.H{"error": "Format JSON salah"})
			return
		}

		db.Create(&UserBaru)
		c.JSON(200, gin.H{"status": "User berhasil ditambahkan"})
	})
}

func GetAllUsers(c *gin.Context) {
	r.GET("/users", func(c *gin.Context) {
		var AllUsers []User

		db.Find(&AllUsers)
		c.JSON(200, gin.H{"data": AllUsers})
	})
}
