package repository

import (
	"github.com/FLA-Official/cars-api/internal/db"
	"github.com/FLA-Official/cars-api/internal/models"
)

// create a new car
func CreateCar(car *models.Car) error {
	result := db.DB.Create(car)
	return result.Error
}

// get a car by ID
func GetCarByID(id uint) (*models.Car, error) {
	var car models.Car
	result := db.DB.First(&car, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &car, nil
}

// Get all cars
func GetAllcars() ([]models.Car, error) {
	var cars []models.Car
	result := db.DB.Find(&cars)
	return cars, result.Error
}

func UpdateCar(id uint, updateCar *models.Car) error {
	var car models.Car
	if err := db.DB.First(&car, id).Error; err != nil {
		return err
	}

	car.Brand = updateCar.Brand
	car.Model = updateCar.Model
	car.Year = updateCar.Year
	car.Price = updateCar.Price
	car.Engine = updateCar.Engine
	car.HorsePower = updateCar.HorsePower

	return db.DB.Save(&car).Error
}

// Delete a car by ID
func DeleteCar(id uint) error {
	return db.DB.Delete(&models.Car{}, id).Error
}
