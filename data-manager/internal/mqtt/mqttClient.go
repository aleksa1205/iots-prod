package mqtt

import (
	"context"
	"fmt"
	"log"
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
		qos:    cfg.Qos}, nil
}

func (c *SensorMqttClient) Publish(payload []byte) error {
	token := c.client.Publish(c.topic, c.qos, false, payload)
	if token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}
