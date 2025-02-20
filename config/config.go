package config

import (
	"os"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
	AppPort    string
}

func LoadConfig() *Config {
	// Читаем переменные окружения
	return &Config{
		DBHost:     os.Getenv("POSTGRES_DB_HOST"),
		DBPort:     os.Getenv("POSTGRES_DB_PORT"),
		DBUser:     os.Getenv("POSTGRES_USER"),
		DBPassword: os.Getenv("POSTGRES_PASSWORD"),
		DBName:     os.Getenv("POSTGRES_DB"),
		DBSSLMode:  os.Getenv("POSTGRES_DB_SSLMODE"),
		AppPort:    os.Getenv("APP_PORT"),
	}
}
