package main

import (
	"data-manager/internal/config"
	infrastructure "data-manager/internal/repositories"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connectionString := config.GetEnvOrPanic(config.EnvKeys.DatabaseConnectionString)

	db, err := infrastructure.ConnectToDatabase(connectionString)
	if err != nil {
		log.Fatalf("Failed to connect to database: %w", err)
	}

	defer db.Close()
}
