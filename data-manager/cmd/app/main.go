package main

import (
	"data-manager/config"
	application "data-manager/internal/application/common"
	infrastructure "data-manager/internal/infrastructure/persistence"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connectionString := config.GetEnvOrPanic(application.EnvKeys.DatabaseConnectionString)

	db, err := infrastructure.ConnectToDatabase(connectionString)
	if err != nil {
		log.Fatalf("Failed to connect to database: %w", err)
	}

	defer db.Close()
}
