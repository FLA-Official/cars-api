package db

import "github.com/FLA-Official/cars-api/internal/models"

func Seed() {
	DB.Create(&models.Car{
		Brand:      "Toyota",
		Model:      "Corolla",
		Year:       2022,
		Price:      30000,
		Engine:     "Petrol",
		HorsePower: 140,
	})
}
