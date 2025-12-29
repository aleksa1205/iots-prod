package config

type EnvironmentVariableKeys struct {
	MqttBroker       string
	MqttClientId     string
	MqttReceiveTopic string
	MqttGenThreshold string
	UsedThreshold    string
	PublishTopic     string
}

var EnvKeys = EnvironmentVariableKeys{
	MqttBroker:       "MQTT_BROKER",
	MqttClientId:     "MQTT_CLIENT_ID",
	MqttReceiveTopic: "MQTT_RECEIVE_TOPIC",
	PublishTopic:     "MQTT_PUBLISH_TOPIC",
	MqttGenThreshold: "GEN_THRESHOLD",
	UsedThreshold:    "USED_THRESHOLD",
}
