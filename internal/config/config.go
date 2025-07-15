package config

import (
	"log"
	"os"
)

type Config struct {
	DatabaseURL string
	ServerAddr  string
	Env         string
}

func (config *Config) IsDev() bool {
	return config.Env == "dev"
}

func (config *Config) IsProd() bool {
	return config.Env == "prod"
}

func Load() Config {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL is required")
	}

	serverAddr := os.Getenv("SERVER_ADDR")
	if serverAddr == "" {
		serverAddr = ":2137"
	}

	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}

	return Config{
		DatabaseURL: dbURL,
		ServerAddr:  serverAddr,
		Env:         env,
	}
}
