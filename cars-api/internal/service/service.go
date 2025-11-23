package service

import (
	"github.com/FLA-Official/cars-api/internal/repository"
)

type CarService struct {
	repo repository.CarRepository
}
