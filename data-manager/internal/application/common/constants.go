package application

type EnvironmentVariableKeys struct {
	Env                      string
	DatabaseConnectionString string
}

var EnvKeys = EnvironmentVariableKeys{
	Env:                      "ENV",
	DatabaseConnectionString: "DB_CONNECTION_STRING",
}
