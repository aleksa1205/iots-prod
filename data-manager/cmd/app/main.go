package main

import (
	"data-manager/internal/config"
	"data-manager/internal/handlers"
	sensorpb "data-manager/internal/proto"
	infrastructure "data-manager/internal/repositories"
	"data-manager/internal/services"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func InitServer() net.Listener {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	port := os.Getenv(config.EnvKeys.Port)
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Error listening for tcp connections on port %v: %v", port, err)
	}

	return listen
}

func DbConnect() *gorm.DB {
	connectionString := config.GetEnvOrPanic(config.EnvKeys.DatabaseConnectionString)

	db, err := infrastructure.ConnectToDatabase(connectionString)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	// Think about closing now here
	//defer db.Close()
	return db
}

func main() {
	lis := InitServer()
	db := DbConnect()

	repository := infrastructure.NewSensorReadingRepository(db)
	service := services.NewSensorReadingService(repository)
	server := grpc.NewServer()
	handler := handlers.NewSensorReadingHandler(service)
	sensorpb.RegisterSensorReadingServiceServer(server, handler)

	log.Println("Starting server...")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
