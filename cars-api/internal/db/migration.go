package db

import (
	"log"

	"github.com/FLA-Official/cars-api/internal/models"
)

func Migrate() {
	err := DB.AutoMigrate(&models.Car{})
	if err != nil {
		log.Fatal("Migration Failed:", err)
	}
	log.Println("Migration Completed")
}
