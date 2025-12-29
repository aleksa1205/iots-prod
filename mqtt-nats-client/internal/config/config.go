package config

type Config struct {
	MqttBroker   string
	MqttClientId string
	MqttTopic    string
	NatsBroker   string
	NatsSubject  string
}

func LoadConfig() *Config {
	return &Config{
		MqttBroker:   GetEnvOrPanic(EnvKeys.MqttBroker),
		MqttClientId: GetEnvOrPanic(EnvKeys.MqttClientId),
		MqttTopic:    GetEnvOrPanic(EnvKeys.MqttTopic),
		NatsSubject:  GetEnvOrPanic(EnvKeys.NatsSubject),
		NatsBroker:   GetEnvOrPanic(EnvKeys.NatsBroker),
	}
}
