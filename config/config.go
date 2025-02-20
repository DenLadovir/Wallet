package config

import (
	"github.com/joho/godotenv"
	"log"
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
	// Загружаем переменные из файла config.env
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Читаем переменные окружения
	return &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBSSLMode:  os.Getenv("DB_SSLMODE"),
		AppPort:    os.Getenv("APP_PORT"),
	}
}
