package config

type Config struct {
	Broker       string
	ClientId     string
	PublishTopic string
	ReceiveTopic string
	MLaasUrl     string
	NatsBroker   string
	Subject      string
}

func LoadConfig() *Config {
	return &Config{
		Broker:       GetEnvOrPanic(EnvKeys.Broker),
		PublishTopic: GetEnvOrPanic(EnvKeys.PublishTopic),
		ReceiveTopic: GetEnvOrPanic(EnvKeys.ReceiveTopic),
		ClientId:     GetEnvOrPanic(EnvKeys.ClientId),
		MLaasUrl:     GetEnvOrPanic(EnvKeys.MLaaSUrl),
		NatsBroker:   GetEnvOrPanic(EnvKeys.NatsBroker),
		Subject:      GetEnvOrPanic(EnvKeys.Subject),
	}
}
