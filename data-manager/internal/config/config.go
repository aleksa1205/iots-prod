package config

type Config struct {
	DatabaseConnectionString string
	MqttBroker               string
	MqttClientId             string
	MqttTopic                string
}

func LoadConfig() *Config {
	return &Config{
		DatabaseConnectionString: GetEnvOrPanic(EnvKeys.DatabaseConnectionString),
		MqttBroker:               GetEnvOrPanic(EnvKeys.MqttBroker),
		MqttClientId:             GetEnvOrPanic(EnvKeys.MqttClientId),
		MqttTopic:                GetEnvOrPanic(EnvKeys.MqttTopic),
	}
}
