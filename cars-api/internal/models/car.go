package models

type Car struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	Brand      string `json:"brand"`
	Model      string `json:"model"`
	Year       int    `json:"year"`
	Price      int    `json:"price"`
	Engine     string `json:"engine"`
	HorsePower int    `json:"horse_power"`
}
