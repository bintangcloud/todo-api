package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "user=postgres password=bintang444 dbname=todo_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal membuka koneksi GORM:", err)
	}
	fmt.Println("Koneksi GORM berhasil")

	db.AutoMigrate(&User{}, &Todo{})
	fmt.Println("Migrasi tabel berhasil")

	r := gin.Default()

	//Endpoint user
	r.POST("/tambah-user", func(c *gin.Context) {
		var UserBaru User
		if err := c.ShouldBindJSON(&UserBaru); err != nil {
			c.JSON(400, gin.H{"error": "Format JSON salah"})
			return
		}

		db.Create(&UserBaru)
		c.JSON(200, gin.H{"status": "User berhasil ditambahkan"})
	})

	r.GET("/tampilkan-user", func(c *gin.Context) {
		var AllUsers []User

		db.Find(&AllUsers)
		c.JSON(200, gin.H{"data": AllUsers})
	})

	r.PUT("/update-user/:id", func(c *gin.Context) {
		id := c.Param("id")
		var UserLama User

		if err := db.First(&UserLama, id).Error; err != nil {
			c.JSON(404, gin.H{"error": "User tidak ditemukan!"})
			return
		}

		var UserBaru User
		if err := c.ShouldBindJSON(&UserBaru); err != nil {
			c.JSON(400, gin.H{"error": "Format JSON salah"})
			return
		}

		db.Model(&UserLama).Updates(UserBaru)
		c.JSON(200, gin.H{"status": "User berhasil diupdate!"})
	})

	r.DELETE("/user/:id", func(c *gin.Context) {
		id := c.Param("id")

		db.Delete(&User{}, id)
		c.JSON(200, gin.H{"status": "sukses menghapus user id" + id})
	})

	//Endpoint TODO
	r.POST("/tambah-todo", func(c *gin.Context) {
		var TodoBaru Todo
		if err := c.ShouldBindJSON(&TodoBaru); err != nil {
			c.JSON(400, gin.H{"error": "Format JSON salah"})
			return
		}

		db.Create(&TodoBaru)
		c.JSON(200, gin.H{"status": "Todo berhasil ditambahkan"})
	})

	r.Run(":8080")
}
