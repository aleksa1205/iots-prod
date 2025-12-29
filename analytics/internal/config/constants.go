package config

type EnvironmentVariableKeys struct {
	MqttBroker   string
	MqttClientId string
	MqttTopic    string
	MLaaSUrl     string
	NatsBroker   string
	NatsSubject  string
}

var EnvKeys = EnvironmentVariableKeys{
	MqttBroker:   "MQTT_BROKER",
	MqttClientId: "MQTT_CLIENT_ID",
	MqttTopic:    "MQTT_TOPIC",
	MLaaSUrl:     "MLAAS_URL",
	NatsBroker:   "NATS_BROKER",
	NatsSubject:  "NATS_SUBJECT",
}
