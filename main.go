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

	r.POST("/tambah-user", func(c *gin.Context) {
		var UserBaru User
		if err := c.ShouldBindJSON(&UserBaru); err != nil {
			c.JSON(400, gin.H{"error": "Format JSON salah"})
			return
		}

		db.Create(&UserBaru)
		c.JSON(200, gin.H{"status": "User berhasil ditambahkan"})
	})

}
