package config

type EnvironmentVariableKeys struct {
	DatabaseConnectionString string
	MqttBroker               string
	MqttClientId             string
	MqttTopic                string
}

var EnvKeys = EnvironmentVariableKeys{
	DatabaseConnectionString: "DB_CONNECTION_STRING",
	MqttBroker:               "MQTT_BROKER",
	MqttClientId:             "MQTT_CLIENT_ID",
	MqttTopic:                "MQTT_TOPIC",
}
