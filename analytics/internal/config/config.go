package config

type Config struct {
	MqttBroker   string
	MqttClientId string
	MqttTopic    string
	MLaaSUrl     string
	NatsBroker   string
	NatsSubject  string
}

func LoadConfig() *Config {
	return &Config{
		MqttBroker:   GetEnvOrPanic(EnvKeys.MqttBroker),
		MqttTopic:    GetEnvOrPanic(EnvKeys.MqttTopic),
		MqttClientId: GetEnvOrPanic(EnvKeys.MqttClientId),
		MLaaSUrl:     GetEnvOrPanic(EnvKeys.MLaaSUrl),
		NatsBroker:   GetEnvOrPanic(EnvKeys.NatsBroker),
		NatsSubject:  GetEnvOrPanic(EnvKeys.NatsSubject),
	}
}
