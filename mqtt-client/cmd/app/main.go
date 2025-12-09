package main

import (
	"log"
	"mqqt-client/internal/config"
	lmqtt "mqqt-client/internal/mqtt"
	"net/http"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/joho/godotenv"
)

func StartServer() {
	address := config.GetEnvOrPanic(config.EnvKeys.Host) + ":" + config.GetEnvOrPanic(config.EnvKeys.Port)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Listening on " + address)
}

func InitMqttClient() (mqtt.Client, error) {
	broker := config.GetEnvOrPanic(config.EnvKeys.Broker)
	clientId := config.GetEnvOrPanic(config.EnvKeys.ClientId)
	return lmqtt.CreateMQTTClient(broker, clientId)
}

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	go StartServer()

	client, err := InitMqttClient()
	if err != nil {
		log.Fatalf("Failed to connect to MQTT client: %v", err)
	}

	lmqtt.ReceiveMessage(client)

	select {}
}
