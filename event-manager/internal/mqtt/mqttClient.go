package mqtt

import (
	"context"
	"encoding/json"
	"event-manager/internal/dtos"
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type ConfigMqtt struct {
	Broker        string
	ClientId      string
	ReceiveTopic  string
	PublishTopic  string
	Qos           byte
	GenThreshold  float64
	UsedThreshold float64
}

type SensorMqttClient struct {
	client       mqtt.Client
	receiveTopic string
	publishTopic string
	qos          byte
	threshold    dtos.EventThreshold
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
		client:       client,
		publishTopic: cfg.PublishTopic,
		receiveTopic: cfg.ReceiveTopic,
		qos:          cfg.Qos,
		threshold: dtos.EventThreshold{
			GenerateKw: cfg.GenThreshold,
			UsedKw:     cfg.UsedThreshold,
		},
	}, nil
}

func (c *SensorMqttClient) Subscribe() error {
	token := c.client.Subscribe(c.receiveTopic, c.qos, c.handleMessage)

	if token.Wait() && token.Error() != nil {
		return fmt.Errorf("failed to subscribe to topic %s: %w", c.receiveTopic, token.Error())
	}

	log.Printf("Subscribed to topic %s", c.receiveTopic)
	return nil
}

func (c *SensorMqttClient) Publish(payload []byte) error {
	token := c.client.Publish(c.publishTopic, c.qos, false, payload)
	if token.Wait() && token.Error() != nil {
		return token.Error()
	}
	return nil
}

func (c *SensorMqttClient) handleMessage(client mqtt.Client, msg mqtt.Message) {
	reading, err := c.receiveMessage(client, msg)
	if err != nil {
		log.Println(err)
		return
	}

	events, detected := c.detectEvent(reading)
	if !detected {
		return
	}

	for _, event := range events {
		payload, err := json.Marshal(event)
		if err != nil {
			log.Println("Failed to marshal event:", err)
			continue
		}
		if err := c.Publish(payload); err != nil {
			log.Println("Failed to publish event:", err)
		}
	}
}

func (c *SensorMqttClient) receiveMessage(_ mqtt.Client, msg mqtt.Message) (dtos.SensorReadingOverview, error) {
	var reading dtos.SensorReadingOverview
	err := json.Unmarshal(msg.Payload(), &reading)
	if err != nil {
		return dtos.SensorReadingOverview{}, fmt.Errorf("error unmarshalling message: topic: %s payload: %s, err: %v", msg.Topic(), msg.Payload(), err)

	}
	log.Printf("%+v", reading)
	return reading, nil
}

func (c *SensorMqttClient) detectEvent(reading dtos.SensorReadingOverview) ([]dtos.AlertEvent, bool) {
	events := make([]dtos.AlertEvent, 0, 2)

	if reading.GeneratedKW > c.threshold.GenerateKw {
		events = append(events, dtos.CreateSensorReadingAlert(reading, dtos.GenerateOverflow))
	}

	if reading.UsedKW > c.threshold.UsedKw {
		events = append(events, dtos.CreateSensorReadingAlert(reading, dtos.UsedOverflow))
	}

	if len(events) == 0 {
		return nil, false
	}

	return events, true
}
