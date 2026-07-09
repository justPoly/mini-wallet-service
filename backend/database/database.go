package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/justPoly/mini-wallet-service/backend/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found")
	}

	// Read database name from .env
	dbName := os.Getenv("DB_NAME")

	// Connect to SQLite
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = db

	fmt.Println("✅ Connected to SQLite database")

	// Automatically create/update database tables
	err = DB.AutoMigrate(
		&models.Account{},
	)

	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	fmt.Println("✅ Database migrated successfully")
}