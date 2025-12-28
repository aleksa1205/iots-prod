package config

type EnvironmentVariableKeys struct {
	Env        string
	MqttBroker string
	ClientId   string
	Topic      string
	NatsBroker string
	Subject    string
}

var EnvKeys = EnvironmentVariableKeys{
	Env:        "ENV",
	MqttBroker: "BROKER",
	ClientId:   "CLIENT_ID",
	Topic:      "TOPIC",
	NatsBroker: "NATS_BROKER",
	Subject:    "SUBJECT",
}
