package main

import (
	"fmt"
	"log"

	"github.com/FLA-Official/cars-api/internal/db"
)

func main() {
	db.Init() // connecting through postgres database
	fmt.Println("Connecting to DB...")

	db.Migrate()
	fmt.Println("Migrating...")

	db.Seed()
	fmt.Println("Seeding...")

	log.Println("Started")
	fmt.Println("Done!")
}
