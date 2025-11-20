package db

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(sqlite.Open("cars.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect DB:", err)
	}

	log.Println("Database Connected and Migrated")
	DB = db
}
