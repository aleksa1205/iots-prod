package config

type EnvironmentVariableKeys struct {
	Env          string
	Broker       string
	ClientId     string
	PublishTopic string
	ReceiveTopic string
	MLaaSUrl     string
	NatsBroker   string
	Subject      string
}

var EnvKeys = EnvironmentVariableKeys{
	Env:          "ENV",
	Broker:       "BROKER",
	ClientId:     "CLIENT_ID",
	PublishTopic: "PUBLISH_TOPIC",
	ReceiveTopic: "RECEIVE_TOPIC",
	MLaaSUrl:     "MLAASURL",
	NatsBroker:   "NATS_BROKER",
	Subject:      "SUBJECT",
}
