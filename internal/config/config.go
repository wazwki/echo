package config

import "os"

type Config struct {
	Host string
	Port string
}

func LoadFromEnv() (*Config, error) {
	cfg := &Config{
		Host: os.Getenv("HOST"),
		Port: os.Getenv("PORT"),
	}

	return cfg, nil
}
