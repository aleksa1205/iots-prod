package config

type EnvironmentVariableKeys struct {
	DatabaseConnectionString string
	Broker                   string
	ClientId                 string
	Topic                    string
}

var EnvKeys = EnvironmentVariableKeys{
	DatabaseConnectionString: "DB_CONNECTION_STRING",
	Broker:                   "BROKER",
	ClientId:                 "CLIENT_ID",
	Topic:                    "TOPIC",
}
