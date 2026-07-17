package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() {
	var DB *gorm.DB

	dsn := "user=postgres password=bintang444 dbname=todo_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal membuka koneksi GORM:", err)
	}
	fmt.Println("Koneksi GORM berhasil")

	db.AutoMigrate(&models.User{}, &models.Todo{})
	fmt.Println("Migrasi tabel berhasil")

	DB = db
}
