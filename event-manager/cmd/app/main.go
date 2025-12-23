package main

import (
	"context"
	"event-manager/internal/config"
	lmqtt "event-manager/internal/mqtt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
)

func loadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Println("No .env file found, using system environment variables")
	}
}

func main() {
	loadEnv()
	cfg := config.LoadConfig()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-signalChan
		log.Println("Terminal signal received")
		cancel()
	}()

	client, err := lmqtt.CreateMQTTClient(ctx, &lmqtt.ConfigMqtt{
		Broker:        cfg.Broker,
		ClientId:      cfg.ClientId,
		GenThreshold:  cfg.GenThreshold,
		UsedThreshold: cfg.UsedThreshold,
		Qos:           1,
		ReceiveTopic:  cfg.ReceiveTopic,
		PublishTopic:  cfg.PublishTopic,
	})
	if err != nil {
		log.Fatal(err)
	}

	err = client.Subscribe()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Analytics service started ...")

	<-ctx.Done()
	log.Println("Shutting down...")
}
