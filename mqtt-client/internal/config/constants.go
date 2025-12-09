package config

type EnvironmentVariableKeys struct {
	Env      string
	Host     string
	Port     string
	Broker   string
	ClientId string
	Topic    string
}

var EnvKeys = EnvironmentVariableKeys{
	Env:      "ENV",
	Host:     "HOST",
	Port:     "PORT",
	Broker:   "BROKER",
	ClientId: "CLIENT_ID",
	Topic:    "TOPIC",
}
