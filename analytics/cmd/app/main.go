package main

import (
	"analytics/internal/config"
	lmqtt "analytics/internal/mqtt"
	"analytics/internal/nats"
	"context"
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

	natsClient, err := nats.CreateNatsClient(ctx, &nats.ConfigNats{
		Broker:  cfg.NatsBroker,
		Subject: cfg.Subject,
	})
	if err != nil {
		log.Fatal(err)
	}

	client, err := lmqtt.CreateMQTTClient(ctx, &lmqtt.ConfigMqtt{
		Broker:       cfg.Broker,
		ClientId:     cfg.ClientId,
		ReceiveTopic: cfg.ReceiveTopic,
		Qos:          1,
		MlaaSUrl:     cfg.MLaasUrl,
		BufferSize:   20,
		NatsClient:   natsClient,
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
