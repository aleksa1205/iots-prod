package mqtt

import (
	"analytics/internal/dtos"
	nats "analytics/internal/nats"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type ConfigMqtt struct {
	Broker       string
	ClientId     string
	ReceiveTopic string
	Qos          byte
	MlaaSUrl     string
	BufferSize   int
	NatsClient   *nats.SensorNatsClient
}

type SensorMqttClient struct {
	client       mqtt.Client
	receiveTopic string
	qos          byte
	buffer       []float64
	bufferSize   int
	mutex        sync.Mutex
	mlaasUrl     string
	natsClient   *nats.SensorNatsClient
}

func CreateMQTTClient(ctx context.Context, cfg *ConfigMqtt) (*SensorMqttClient, error) {
	broker := cfg.Broker
	clientId := cfg.ClientId
	receiveTopic := cfg.ReceiveTopic
	qos := cfg.Qos

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
		receiveTopic: receiveTopic,
		qos:          qos,
		mlaasUrl:     cfg.MlaaSUrl,
		bufferSize:   cfg.BufferSize,
		natsClient:   cfg.NatsClient,
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

func (c *SensorMqttClient) handleMessage(client mqtt.Client, msg mqtt.Message) {
	reading, err := c.receiveMessage(client, msg)
	if err != nil {
		log.Println(err)
		return
	}

	c.mutex.Lock()
	c.buffer = append(c.buffer, reading.UsedKW)
	if len(c.buffer) > c.bufferSize {
		batch := c.buffer[:c.bufferSize]
		c.buffer = c.buffer[c.bufferSize:]
		go c.sendToMLaaS(batch)
	}
	c.mutex.Unlock()
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

func (c *SensorMqttClient) sendToMLaaS(batch []float64) {
	reqBody := map[string]interface{}{
		"past_values": batch, // just UsedKW values
	}

	data, err := json.Marshal(reqBody)
	if err != nil {
		log.Println("MLaaS request failed: ", err)
		return
	}

	resp, err := http.Post(c.mlaasUrl+"/predict", "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Println("MLaaS request failed: ", err)
		return
	}
	defer resp.Body.Close()

	var result struct {
		Prediction float64 `json:"prediction"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Println("Failed to decode MLaaS response:", err)
		return
	}
	log.Printf("Analytics prediction for batch: %f", result.Prediction)

	if c.natsClient != nil {
		data, err := json.Marshal(&nats.AnalyticsResult{
			Prediction: result.Prediction,
			Timestamp:  time.Now().Unix(),
			Model:      "linear-regression",
		})
		if err != nil {
			log.Println("Failed to marshal analytics result:", err)
		}
		
		err = c.natsClient.Publish(data)
		if err != nil {
			log.Println("Failed to publish analytics prediction:", err)
		}
	}
}
