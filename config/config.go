package config

import (
	"os"
	"strconv"
)

type Config struct {
	Port int
}

var AppConfig Config

func loadConfig() {
	AppConfig = Config{
		Port: 8000,
	}

	// Load port from environment variable
	if port, exists := os.LookupEnv("PORT"); exists {
		if portInt, err := strconv.Atoi(port); err == nil {
			AppConfig.Port = portInt
		}
	}
}

func init() {
	loadConfig()
}
