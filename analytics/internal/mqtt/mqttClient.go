package mqtt

import (
	"analytics/internal/config"
	"analytics/internal/dtos"
	"encoding/json"
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func CreateMQTTClient(broker string, clientId string) (mqtt.Client, error) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(clientId)
	opts.SetKeepAlive(60 * time.Second)
	opts.SetPingTimeout(1 * time.Second)
	opts.SetAutoReconnect(true)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, fmt.Errorf("Failed to connect to MQTT broker %s: %w", broker, token.Error())
	}

	log.Printf("Connect to the MQTT broker")
	return client, nil
}

func ReceiveMessage(client mqtt.Client) {
	topic := config.GetEnvOrPanic(config.EnvKeys.Topic)
	client.Subscribe(topic, 1, func(client mqtt.Client, msg mqtt.Message) {
		var reading dtos.SensorReadingOverview
		err := json.Unmarshal(msg.Payload(), &reading)
		if err != nil {
			log.Printf("Error unmarshalling received message: %s", err)
		}
		log.Printf("Received data from topic %s: %+v", topic, reading)
	})
}
