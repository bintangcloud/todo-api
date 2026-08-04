package database

import (
	"fmt"
	"log"
	"os"
	"todo-api/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Gagal membuka koneksi GORM:", err)
	}

	fmt.Println("Koneksi GORM berhasil")

	db.AutoMigrate(&models.User{}, &models.Todo{})

	fmt.Println("Migrasi tabel berhasil")

	DB = db
}
