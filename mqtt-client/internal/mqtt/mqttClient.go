package mqtt

import (
	"encoding/json"
	"fmt"
	"log"
	"mqqt-client/internal/config"
	"mqqt-client/internal/dtos"
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
	return client, nil
}

func ReceiveMessage(client mqtt.Client) {
	topic := config.GetEnvOrPanic(config.EnvKeys.Topic)
	client.Subscribe(topic, 1, func(client mqtt.Client, msg mqtt.Message) {
		var event dtos.AlertEvent
		err := json.Unmarshal(msg.Payload(), &event)
		if err != nil {
			log.Printf("Error unmarshalling event: %v", err)
		}
		log.Printf("Received data from topic %s: %+v", topic, event)
	})
}
