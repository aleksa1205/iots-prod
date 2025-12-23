package config

type EnvironmentVariableKeys struct {
	Env      string
	Broker   string
	ClientId string
	Topic    string
}

var EnvKeys = EnvironmentVariableKeys{
	Env:      "ENV",
	Broker:   "BROKER",
	ClientId: "CLIENT_ID",
	Topic:    "TOPIC",
}
