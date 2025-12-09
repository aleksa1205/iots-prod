package mqtt

import (
	"encoding/json"
	"event-manager/internal/config"
	"event-manager/internal/dtos"
	"fmt"
	"log"
	"strconv"
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
	genThreshold, err := strconv.ParseFloat(config.GetEnvOrPanic(config.EnvKeys.GenThreshold), 64)
	if err != nil {
		log.Printf("Invalid %s using default: %v", config.EnvKeys.GenThreshold, err)
		genThreshold = 0
	}

	usedThreshold, err := strconv.ParseFloat(config.GetEnvOrPanic(config.EnvKeys.UsedThreshold), 64)
	if err != nil {
		log.Printf("Invalid %s using default: %v", config.EnvKeys.GenThreshold, err)
		usedThreshold = 0
	}

	recvTopic := config.GetEnvOrPanic(config.EnvKeys.ReceiveTopic)
	sendTopic := config.GetEnvOrPanic(config.EnvKeys.SendTopic)
	client.Subscribe(recvTopic, 1, func(client mqtt.Client, msg mqtt.Message) {
		var reading dtos.SensorReadingOverview
		err := json.Unmarshal(msg.Payload(), &reading)
		if err != nil {
			log.Printf("Error unmarshalling received data: %s", err)
		}
		log.Printf("Received data from topic %s: %+v", recvTopic, reading)

		if reading.GeneratedKW > genThreshold || reading.UsedKW > usedThreshold {
			alert := dtos.CreateSensorReadingAlert(reading)
			payload, err := json.Marshal(&alert)
			if err != nil {
				log.Printf("Error marshalling received data: %s", err)
			}
			err = publishToTopic(client, sendTopic, payload)
			if err != nil {
				log.Printf("Error publishing to topic %s: %v", sendTopic, err)
			}
		}
	})
}

func publishToTopic(client mqtt.Client, topic string, payload []byte) error {
	token := client.Publish(topic, 1, false, payload)
	if token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}
