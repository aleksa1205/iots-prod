package config

type EnvironmentVariableKeys struct {
	Env                      string
	DatabaseConnectionString string
	Port                     string
	Host                     string
	Broker                   string
	ClientId                 string
}

var EnvKeys = EnvironmentVariableKeys{
	Env:                      "ENV",
	DatabaseConnectionString: "DB_CONNECTION_STRING",
	Host:                     "HOST",
	Port:                     "PORT",
	Broker:                   "BROKER",
	ClientId:                 "CLIENT_ID",
}
