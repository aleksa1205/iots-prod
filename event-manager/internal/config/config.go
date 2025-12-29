package config

type Config struct {
	MqttBroker       string
	MqttClientId     string
	MqttReceiveTopic string
	MqttPublishTopic string
	UsedThreshold    float64
	GenThreshold     float64
}

func LoadConfig() *Config {
	return &Config{
		MqttBroker:       GetEnvOrPanic(EnvKeys.MqttBroker),
		MqttClientId:     GetEnvOrPanic(EnvKeys.MqttClientId),
		MqttReceiveTopic: GetEnvOrPanic(EnvKeys.MqttReceiveTopic),
		MqttPublishTopic: GetEnvOrPanic(EnvKeys.PublishTopic),
		UsedThreshold:    GetEnvFloatOrPanic(EnvKeys.UsedThreshold),
		GenThreshold:     GetEnvFloatOrPanic(EnvKeys.MqttGenThreshold),
	}
}
