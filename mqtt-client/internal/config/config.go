package config

type Config struct {
	Broker   string
	ClientId string
	Topic    string
}

func LoadConfig() *Config {
	return &Config{
		Broker:   GetEnvOrPanic(EnvKeys.Broker),
		ClientId: GetEnvOrPanic(EnvKeys.ClientId),
		Topic:    GetEnvOrPanic(EnvKeys.Topic),
	}
}
