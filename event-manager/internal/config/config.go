package config

type Config struct {
	Broker        string
	ClientId      string
	ReceiveTopic  string
	PublishTopic  string
	UsedThreshold float64
	GenThreshold  float64
}

func LoadConfig() *Config {
	return &Config{
		Broker:        GetEnvOrPanic(EnvKeys.Broker),
		ClientId:      GetEnvOrPanic(EnvKeys.ClientId),
		ReceiveTopic:  GetEnvOrPanic(EnvKeys.ReceiveTopic),
		PublishTopic:  GetEnvOrPanic(EnvKeys.PublishTopic),
		UsedThreshold: GetEnvFloatOrPanic(EnvKeys.UsedThreshold),
		GenThreshold:  GetEnvFloatOrPanic(EnvKeys.GenThreshold),
	}
}
