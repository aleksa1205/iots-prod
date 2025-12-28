package nats

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

type AnalyticsResult struct {
	Prediction float64
	Timestamp  int64
	Model      string
}

type ConfigNats struct {
	Broker  string
	Subject string
}

type SensorNatsClient struct {
	client  *nats.Conn
	subject string
}

func CreateNatsClient(ctx context.Context, cfg *ConfigNats) (*SensorNatsClient, error) {
	opts := []nats.Option{
		nats.ReconnectWait(2 * time.Second),
		nats.DisconnectErrHandler(func(nc *nats.Conn, err error) {
			log.Printf("Disconnected from NATS: %v", err)
		}),
		nats.ReconnectHandler(func(nc *nats.Conn) {
			log.Printf("Reconnected to NATS at %s", nc.ConnectedUrl())
		}),
		nats.ClosedHandler(func(nc *nats.Conn) {
			log.Println("NATS connection closed")
		}),
	}

	nc, err := nats.Connect(cfg.Broker, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to NATS %s: %w", cfg.Broker, err)
	}

	log.Printf("Connected to NATS at %s", cfg.Broker)

	go func() {
		<-ctx.Done()
		log.Println("Context canceled — closing NATS connection")
		nc.Close()
	}()

	return &SensorNatsClient{
		client:  nc,
		subject: cfg.Subject,
	}, nil
}

func (n *SensorNatsClient) Publish(payload []byte) error {
	return n.client.Publish(n.subject, payload)
}
