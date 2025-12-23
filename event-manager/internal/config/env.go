package config

import (
	"fmt"
	"os"
	"strconv"
)

func GetEnvOrPanic(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Sprintf("Configuration Error: Missing environment variable %s", key))
	}
	return value
}

func GetEnvFloatOrPanic(key string) float64 {
	value := GetEnvOrPanic(key)

	f, err := strconv.ParseFloat(value, 64)
	if err != nil {
		panic(fmt.Sprintf("invalid float value for %s: %v", key, err))
	}

	return f
}
