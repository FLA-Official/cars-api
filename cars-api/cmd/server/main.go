package main

import (
	"fmt"
	"log"

	"github.com/FLA-Official/cars-api/internal/db"
)

func main() {
	database := db.Init()
	db.Connect()
	fmt.Println("Connecting to DB...")
	db.Migrate()
	fmt.Println("Migrating...")
	db.Seed()
	fmt.Println("Seeding...")
	log.Println("Started")
	fmt.Println("Done!")
}
