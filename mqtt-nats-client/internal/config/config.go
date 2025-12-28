package config

type Config struct {
	MqttBroker string
	ClientId   string
	Topic      string
	NatsBroker string
	Subject    string
}

func LoadConfig() *Config {
	return &Config{
		MqttBroker: GetEnvOrPanic(EnvKeys.MqttBroker),
		ClientId:   GetEnvOrPanic(EnvKeys.ClientId),
		Topic:      GetEnvOrPanic(EnvKeys.Topic),
		Subject:    GetEnvOrPanic(EnvKeys.Subject),
		NatsBroker: GetEnvOrPanic(EnvKeys.NatsBroker),
	}
}
