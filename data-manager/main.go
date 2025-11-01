package main

import (
	"data-manager/internal/infrastructure/persistence"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./config/.env")
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	err = persistence.ConnectToDatabase()
	if err != nil {
		log.Fatalln("Failed to connect to database: %w", err)
	}
	log.Printf("munem")
}
