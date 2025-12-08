package config

type EnvironmentVariableKeys struct {
	Env           string
	Host          string
	Port          string
	Broker        string
	ClientId      string
	ReceiveTopic  string
	GenThreshold  string
	UsedThreshold string
	SendTopic     string
}

var EnvKeys = EnvironmentVariableKeys{
	Env:           "ENV",
	Host:          "HOST",
	Port:          "PORT",
	Broker:        "BROKER",
	ClientId:      "CLIENT_ID",
	ReceiveTopic:  "RECV_TOPIC",
	SendTopic:     "SEND_TOPIC",
	GenThreshold:  "GEN_THRESHOLD",
	UsedThreshold: "USED_THRESHOLD",
}
