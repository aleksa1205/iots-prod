package config

type EnvironmentVariableKeys struct {
	Env                      string
	DatabaseConnectionString string
	Port                     string
	Host                     string
}

var EnvKeys = EnvironmentVariableKeys{
	Env:                      "ENV",
	DatabaseConnectionString: "DB_CONNECTION_STRING",
	Host:                     "HOST",
	Port:                     "PORT",
}
