package main

import (
	"data-manager/internal/config"
	"data-manager/internal/handlers"
	lmqtt "data-manager/internal/mqtt"
	sensorpb "data-manager/internal/proto"
	infrastructure "data-manager/internal/repositories"
	"data-manager/internal/services"
	"database/sql"
	"log"
	"net"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func InitServer() net.Listener {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	address := config.GetEnvOrPanic(config.EnvKeys.Host) + ":" + config.GetEnvOrPanic(config.EnvKeys.Port)
	listen, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error listening for tcp connections on %v: %v", address, err)
	}
	log.Println("Listening on " + address)

	return listen
}

func DbConnect() (*gorm.DB, *sql.DB) {
	connectionString := config.GetEnvOrPanic(config.EnvKeys.DatabaseConnectionString)

	db, err := infrastructure.ConnectToDatabase(connectionString)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	sqlDb, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get raw database: %v", err)
	}

	return db, sqlDb
}

func MqqtClient() (mqtt.Client, error) {
	broker := config.GetEnvOrPanic(config.EnvKeys.Broker)
	clientId := config.GetEnvOrPanic(config.EnvKeys.ClientId)

	return lmqtt.CreateMQTTClient(broker, clientId)
}

func main() {
	lis := InitServer()
	db, sqlDb := DbConnect()
	defer sqlDb.Close()

	repository := infrastructure.NewSensorReadingRepository(db)
	service := services.NewSensorReadingService(repository)
	server := grpc.NewServer()
	broker, err := MqqtClient()
	if err != nil {
		log.Fatalf("Failed to connect to mqtt client: %v", err)
	}
	handler := handlers.NewSensorReadingHandler(service, broker)
	sensorpb.RegisterSensorReadingServiceServer(server, handler)

	log.Println("Starting server...")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
