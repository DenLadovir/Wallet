package testDB

import (
	"Wallet/database"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
)

var DB *sql.DB

func ClearTestDB() {
	if database.DB == nil {
		log.Fatalf("Database connection is not initialized")
	}

	_, err := database.DB.Exec("TRUNCATE TABLE wallets RESTART IDENTITY CASCADE")
	if err != nil {
		log.Fatalf("Failed to clear test database: %v", err)
	}
	fmt.Println("Test database cleared successfully!")
}

func SeedTestDB(data []map[string]interface{}) {
	for _, wallet := range data {
		_, err := DB.Exec(`
			INSERT INTO wallets (id, balance) VALUES ($1, $2)
		`, wallet["id"], wallet["balance"])
		if err != nil {
			log.Fatalf("Failed to seed test database: %v", err)
		}
	}
}

func InitTestDB() {
	fmt.Println("Initializing test database...")

	// Получаем имя тестовой базы данных
	testDBName := os.Getenv("DB_NAME")
	if testDBName == "" {
		log.Fatalf("Environment variable DB_NAME is not set")
	}

	// Формируем строку подключения к основной базе данных
	mainConnStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s sslmode=%s dbname=postgres",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_SSLMODE"),
	)
	fmt.Println("Main connection string:", mainConnStr)

	// Подключаемся к основной базе данных
	mainDB, err := sql.Open("postgres", mainConnStr)
	if err != nil {
		log.Fatalf("Failed to connect to PostgreSQL: %v", err)
	}
	defer mainDB.Close()

	// Проверяем, существует ли база данных, и создаём её, если она отсутствует
	_, err = mainDB.Exec(fmt.Sprintf("CREATE DATABASE %s", testDBName))
	if err != nil && !strings.Contains(err.Error(), "already exists") {
		log.Fatalf("Failed to create test database: %v", err)
	}

	// Формируем строку подключения к тестовой базе данных
	testConnStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s sslmode=%s dbname=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_SSLMODE"),
		testDBName,
	)
	fmt.Println("Test connection string:", testConnStr)

	// Подключаемся к тестовой базе данных
	database.DB, err = sql.Open("postgres", testConnStr)
	if err != nil {
		log.Fatalf("Failed to connect to test database: %v", err)
	}

	// Проверка соединения
	err = database.DB.Ping()
	if err != nil {
		log.Fatalf("Failed to ping test database: %v", err)
	}
	fmt.Println("Test database connected successfully!")

	// Создаём таблицы и добавляем данные
	_, err = database.DB.Exec(`
		CREATE TABLE IF NOT EXISTS wallets (
			id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
			balance NUMERIC(15, 2) NOT NULL DEFAULT 0.00,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		log.Fatalf("Failed to create table 'wallets': %v", err)
	}

	_, err = database.DB.Exec(`
		INSERT INTO wallets (id, balance) VALUES
			('123e4567-e89b-12d3-a456-426614174000', 1000.00),
			('223e4567-e89b-12d3-a456-426614174001', 500.00),
			('323e4567-e89b-12d3-a456-426614174002', 0.00)
		ON CONFLICT DO NOTHING
	`)
	if err != nil {
		log.Fatalf("Failed to insert initial data into 'wallets': %v", err)
	}
	fmt.Println("Test database initialized successfully!")
}

// checkEnvVars проверяет наличие всех необходимых переменных окружения
//func checkEnvVars() {
//	requiredVars := []string{"DB_HOST", "DB_PORT", "DB_USER", "DB_PASSWORD", "DB_SSLMODE", "DB_TEST_NAME"}
//	for _, v := range requiredVars {
//		if os.Getenv(v) == "" {
//			log.Fatalf("Environment variable %s is not set", v)
//		}
//	}
//}
