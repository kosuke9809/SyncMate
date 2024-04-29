package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDBWithRetry(maxAttempts int, initialBackoff time.Duration) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("Successfully connected to the database")
			return db, nil
		}

		log.Printf("Failed to connect to database. Attempt %d/%d: %v", attempt, maxAttempts, err)
		time.Sleep(initialBackoff * time.Duration(attempt*attempt))
	}
	log.Println("Failed to connect to database after all attempts")
	return nil, err
}

func CloseDB(db *gorm.DB) {
	DB, err := db.DB()
	if err != nil {
		log.Println("Error getting generic database object from GORM:", err)
		return
	}

	if err := DB.Close(); err != nil {
		log.Println("Error closing database connection:", err)
	}
}
