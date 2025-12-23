package mqtt

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"mqqt-client/internal/dtos"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type ConfigMqtt struct {
	Broker   string
	ClientId string
	Topic    string
	Qos      byte
}

type SensorMqttClient struct {
	client mqtt.Client
	topic  string
	qos    byte
}

func CreateMQTTClient(ctx context.Context, cfg *ConfigMqtt) (*SensorMqttClient, error) {
	broker := cfg.Broker
	clientId := cfg.ClientId

	opts := mqtt.NewClientOptions()
	opts.AddBroker(broker)
	opts.SetClientID(clientId)
	opts.SetKeepAlive(60 * time.Second)
	opts.SetPingTimeout(1 * time.Second)
	opts.SetAutoReconnect(true)
	opts.OnConnect = func(client mqtt.Client) {
		log.Printf("Connected to MQTT broker %s", broker)
	}
	opts.OnConnectionLost = func(client mqtt.Client, err error) {
		log.Printf("Disconnected from MQTT broker %s %v", broker, err)
	}

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		return nil, fmt.Errorf("failed to connect to MQTT broker %s: %w", broker, token.Error())
	}

	go func() {
		<-ctx.Done()
		if client.IsConnected() {
			log.Println("Context canceled — disconnecting MQTT client")
			client.Disconnect(250)
		}
	}()

	return &SensorMqttClient{
		client: client,
		topic:  cfg.Topic,
		qos:    cfg.Qos,
	}, nil
}

func (c *SensorMqttClient) Subscribe() error {
	token := c.client.Subscribe(c.topic, c.qos, c.receiveMessage)

	if token.Wait() && token.Error() != nil {
		return fmt.Errorf("failed to subscribe to topic %s: %w", c.topic, token.Error())
	}

	log.Printf("Subscribed to topic %s", c.topic)
	return nil
}

func (c *SensorMqttClient) receiveMessage(_ mqtt.Client, msg mqtt.Message) {
	var reading dtos.AlertEvent
	err := json.Unmarshal(msg.Payload(), &reading)
	if err != nil {
		log.Printf("Error unmarshalling message: topic: %s payload: %s, err: %v", msg.Topic(), msg.Payload(), err)
	}
	log.Printf("%+v", reading)
}
