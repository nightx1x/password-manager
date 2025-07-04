package config

import "os"

type Config struct {
	StoragePath string
}

func NewConfig() *Config {
	path := os.Getenv("PASSWORD_STORE")
	if path == "" {
		path = "passwords.json"
	}
	return &Config{
		StoragePath: path,
	}
}
