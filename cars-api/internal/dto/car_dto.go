package dto

type CreateCarRequest struct {
	Brand string `json:"brand" validate:"required"`
	Model string `json:"model" validate:"required"`
	Year  int    `json:"year" validate:"required,min=1885"`
	Price int    `json:"price" validate:"required"`
}

type UpdateCarRequest struct {
	Brand string `json:"brand"`
	Model string `json:"model"`
	Year  int    `json:"year"`
	Price int    `json:"price"`
}

type CarResponse struct {
	ID    uint   `json:"id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Year  int    `json:"year"`
	Price int    `json:"price"`
}
