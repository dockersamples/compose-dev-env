package config

import (
	"os"
)

type Config struct {
	DB struct {
		User     string
		Password string
		Host     string
		Port     string
		Name     string
	}
	Server struct {
		Port string
	}
}

func Load() *Config {
	cfg := &Config{}
	// Load config from environment variables or config file
	cfg.DB.User = os.Getenv("DB_USER")
	cfg.DB.Password = os.Getenv("DB_PASSWORD")
	cfg.DB.Name = os.Getenv("DB_NAME")
	cfg.DB.Host = os.Getenv("DB_HOST")
	portStr := os.Getenv("DB_PORT")
	if portStr == "" {
		portStr = "3306"
	}
	cfg.DB.Port = portStr

	cfg.Server.Port = "8000"

	return cfg
}
