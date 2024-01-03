package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port int
    Protocol string
}

func New() *Config {
	return &Config{
		Port: getEnvInt("PORT", 6379),
        Protocol: getEnv("PROTOCOL", "tcp"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvInt(key string, fallback int) int {
	strValue := getEnv(key, "")
	intValue, err := strconv.Atoi(strValue)
	if err == nil {
		return intValue
	}

	return fallback
}
