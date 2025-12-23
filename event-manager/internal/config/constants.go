package config

type EnvironmentVariableKeys struct {
	Env           string
	Broker        string
	ClientId      string
	ReceiveTopic  string
	GenThreshold  string
	UsedThreshold string
	PublishTopic  string
}

var EnvKeys = EnvironmentVariableKeys{
	Env:           "ENV",
	Broker:        "BROKER",
	ClientId:      "CLIENT_ID",
	ReceiveTopic:  "RECEIVE_TOPIC",
	PublishTopic:  "PUBLISH_TOPIC",
	GenThreshold:  "GEN_THRESHOLD",
	UsedThreshold: "USED_THRESHOLD",
}
