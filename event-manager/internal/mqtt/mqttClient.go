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

	receiveMessage(client)

	fmt.Println("Connect to the MQTT broker")
	return client, nil
}

func receiveMessage(client mqtt.Client) {
	gen_threshold, err := strconv.ParseFloat(config.GetEnvOrPanic(config.EnvKeys.GenThreshold), 64)
	if err != nil {
		log.Printf("Invalid %s using default: %v", config.EnvKeys.GenThreshold, err)
		gen_threshold = 0
	}

	used_threshold, err := strconv.ParseFloat(config.GetEnvOrPanic(config.EnvKeys.UsedThreshold), 64)
	if err != nil {
		log.Printf("Invalid %s using default: %v", config.EnvKeys.GenThreshold, err)
		used_threshold = 0
	}

	recv_topic := config.GetEnvOrPanic(config.EnvKeys.ReceiveTopic)
	send_topic := config.GetEnvOrPanic(config.EnvKeys.SendTopic)
	client.Subscribe(recv_topic, 1, func(client mqtt.Client, msg mqtt.Message) {
		var reading dtos.SensorReadingOverview
		err := json.Unmarshal(msg.Payload(), &reading)
		if err != nil {
			log.Printf("Error unmarshalling received data: %s", err)
		}
		fmt.Println("Received data from topic %s: %v", recv_topic, reading)

		if reading.GeneratedKW > gen_threshold || reading.UsedKW > used_threshold {
			err := publishToTopic(client, send_topic, reading)
			if err != nil {
				log.Printf("Error publishing to topic %s: %v", send_topic, err)
			}
		}
	})
}

func publishToTopic(client mqtt.Client, topic string, event any) error {
	payload, err := json.Marshal(event)
	if err != nil {
		log.Printf("Error marshalling event: %s", err)
		return err
	}

	token := client.Publish(topic, 1, false, payload)
	if token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}
