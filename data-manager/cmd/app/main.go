package main

import (
	"data-manager/internal/config"
	"data-manager/internal/entities"
	"data-manager/internal/handlers"
	lmqtt "data-manager/internal/mqtt"
	sensorpb "data-manager/internal/proto"
	"data-manager/internal/repositories"
	db2 "data-manager/internal/repositories/db"
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

func InitDb() (*gorm.DB, *sql.DB) {
	connectionString := config.GetEnvOrPanic(config.EnvKeys.DatabaseConnectionString)

	db, err := db2.Connect(connectionString)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	sqlDb, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get database handle: %v", err)
	}

	if err := db.AutoMigrate(&entities.SensorReading{}); err != nil {
		log.Fatalf("Migration failed: %v", err)
	}

	return db, sqlDb
}

func InitMqttClient() (mqtt.Client, error) {
	broker := config.GetEnvOrPanic(config.EnvKeys.Broker)
	clientId := config.GetEnvOrPanic(config.EnvKeys.ClientId)
	return lmqtt.CreateMQTTClient(broker, clientId)
}

func main() {
	lis := InitServer()
	db, sqlDb := InitDb()
	defer sqlDb.Close()

	repository := repositories.NewSensorReadingRepository(db)

	broker, err := InitMqttClient()
	if err != nil {
		log.Fatalf("Failed to connect to MQTT client: %v", err)
	}
	topic := config.GetEnvOrPanic(config.EnvKeys.Topic)
	service := services.NewSensorReadingService(repository, broker, topic)

	server := grpc.NewServer()

	handler := handlers.NewSensorReadingHandler(service)
	sensorpb.RegisterSensorReadingServiceServer(server, handler)

	log.Println("Starting server...")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
