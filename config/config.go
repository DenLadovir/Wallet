package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DBHost     string `envconfig:"POSTGRES_DB_HOST"`
	DBPort     string `envconfig:"POSTGRES_DB_PORT"`
	DBUser     string `envconfig:"POSTGRES_USER"`
	DBPassword string `envconfig:"POSTGRES_PASSWORD"`
	DBName     string `envconfig:"POSTGRES_DB"`
	DBSSLMode  string `envconfig:"POSTGRES_DB_SSLMODE"`
	AppPort    string `envconfig:"APP_PORT"`
}

func LoadConfig() (*Config, error) {
	var cfg Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
