package main

import (
	"data-manager/internal/config"
	"data-manager/internal/handlers"
	sensorpb "data-manager/internal/proto"
	infrastructure "data-manager/internal/repositories"
	"data-manager/internal/services"
	"log"
	"net"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	list, err := net.Listen("tcp", ":55555")

	if err != nil {
		log.Fatal(err)
	}

	err = godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connectionString := config.GetEnvOrPanic(config.EnvKeys.DatabaseConnectionString)

	db, err := infrastructure.ConnectToDatabase(connectionString)
	if err != nil {
		log.Fatalf("Failed to connect to database: %w", err)
	}
	defer db.Close()

	grpcServer := grpc.NewServer()
	sensorSevice := services.SensorReadingService{}
	sensorHandler := handlers.NewSensorReadingHandler(sensorSevice)
	sensorpb.RegisterSensorReadingServiceServer(grpcServer, sensorHandler)

	if err := grpcServer.Serve(list); err != nil {
		log.Fatal(err)
	}
}
