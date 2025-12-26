package config

type Config struct {
	DatabaseConnectionString string
	Broker                   string
	ClientId                 string
	Topic                    string
}

func LoadConfig() *Config {
	return &Config{
		DatabaseConnectionString: GetEnvOrPanic(EnvKeys.DatabaseConnectionString),
		Broker:                   GetEnvOrPanic(EnvKeys.Broker),
		ClientId:                 GetEnvOrPanic(EnvKeys.ClientId),
		Topic:                    GetEnvOrPanic(EnvKeys.Topic),
	}
}
