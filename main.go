package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"user-reader-service/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func initDB() {
	var err error
	err = godotenv.Load()
	if err != nil {
		// Log.Fatal akan mencetak pesan dan menghentikan program jika file .env tidak ditemukan
		log.Fatalf("Error loading .env file: %v", err)
	}
	// Logika initDB sama seperti service creator
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	
	// Catatan: Service reader tidak perlu AutoMigrate, tapi tidak ada salahnya
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Gagal terhubung ke database: %v", err)
	}
	log.Println("Koneksi database berhasil.")
}

// GetUsers adalah handler untuk request GET
func GetUsers(c *gin.Context) {
	var users []models.User
	// Ambil semua user dari database
	if err := DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Gagal mengambil daftar user: " + err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, users)
}

func main() {
	initDB()
	PORT := os.Getenv("ENV_PORT")
	r := gin.Default()
	r.GET("/api/users", GetUsers) // Gunakan endpoint yang sama untuk konsistensi

	log.Println("User Reader Service berjalan di port " + PORT)
	r.Run(fmt.Sprintf(":%s", PORT))
}