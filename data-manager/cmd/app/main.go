package main

import (
	"context"
	"data-manager/internal/config"
	"data-manager/internal/entities"
	"data-manager/internal/handlers"
	lmqtt "data-manager/internal/mqtt"
	sensorpb "data-manager/internal/proto"
	"data-manager/internal/repositories"
	ldb "data-manager/internal/repositories/db"
	"data-manager/internal/services"
	"database/sql"
	"errors"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

func loadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found, using system environment variables")
	}
}

func initDb(connectionString string) (*gorm.DB, *sql.DB) {
	db, err := ldb.Connect(connectionString)
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

func main() {
	loadEnv()
	cfg := config.LoadConfig()

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	//lis := InitServer()
	db, sqlDb := initDb(cfg.DatabaseConnectionString)
	defer sqlDb.Close()

	repository := repositories.NewSensorReadingRepository(db)

	client, err := lmqtt.CreateMQTTClient(ctx, &lmqtt.ConfigMqtt{
		Broker:   cfg.MqttBroker,
		Qos:      1,
		Topic:    cfg.MqttTopic,
		ClientId: cfg.MqttClientId})
	if err != nil {
		log.Fatal(err)
	}
	service := services.NewSensorReadingService(repository, client, cfg.MqttTopic)

	server := grpc.NewServer()
	handler := handlers.NewSensorReadingHandler(service)

	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Error listening for tcp connections on :8080, %v", err)
	}

	sensorpb.RegisterSensorReadingServiceServer(server, handler)

	go func() {
		log.Printf("Server listening at %v", listen.Addr())
		if err := server.Serve(listen); err != nil {
			if !errors.Is(err, grpc.ErrServerStopped) {
				log.Printf("GRPC server failed to serve: %v", err)
				stop()
			}
		}
	}()

	<-ctx.Done()
	log.Println("Shutting down server...")
	server.GracefulStop()
}
