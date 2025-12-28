package main

import (
	"context"
	"log"
	"mqqt-client/internal/config"
	lmqtt "mqqt-client/internal/mqtt"
	nats2 "mqqt-client/internal/nats"
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
		Broker:   cfg.MqttBroker,
		ClientId: cfg.ClientId,
		Topic:    cfg.Topic,
		Qos:      1,
	})
	if err != nil {
		log.Fatal(err)
	}

	err = client.Subscribe()
	if err != nil {
		log.Fatal(err)
	}

	nats, err := nats2.CreateNatsClient(ctx, &nats2.ConfigNats{
		Subject: cfg.Subject,
		Broker:  cfg.NatsBroker,
	})
	if err != nil {
		log.Fatal(err)
	}

	err = nats.Subscribe()

	log.Println("Mqtt/Nats Client service started ...")

	<-ctx.Done()
	log.Println("Shutting down...")
}
