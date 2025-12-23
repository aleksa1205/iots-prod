package config

type Config struct {
	Broker       string
	ClientId     string
	PublishTopic string
	ReceiveTopic string
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		Broker:       GetEnvOrPanic(EnvKeys.Broker),
		PublishTopic: GetEnvOrPanic(EnvKeys.PublishTopic),
		ReceiveTopic: GetEnvOrPanic(EnvKeys.ReceiveTopic),
		ClientId:     GetEnvOrPanic(EnvKeys.ClientId),
	}

	return cfg, nil
}
