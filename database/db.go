package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitTestDB() {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to test database: %v", err)
	}

	// Проверка соединения
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Failed to ping test database: %v", err)
	}
	fmt.Println("Database for tests connected successfully!")
}

func InitDB() {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	fmt.Println("Database connected successfully!")
}

//func InitTestDB() {
//	connStr := fmt.Sprintf(
//		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
//		os.Getenv("DB_HOST"),
//		os.Getenv("DB_PORT"),
//		os.Getenv("DB_USER"),
//		os.Getenv("DB_PASSWORD"),
//		os.Getenv("DB_NAME"),
//		os.Getenv("DB_SSLMODE"),
//	)
//
//	var err error
//	// Используем "=" вместо ":=" для обновления глобальной переменной DB
//	DB, err = sql.Open("postgres", connStr)
//	if err != nil {
//		log.Fatalf("Failed to connect to test database: %v", err)
//	}
//
//	// Проверка соединения
//	err = DB.Ping()
//	if err != nil {
//		log.Fatalf("Failed to ping test database: %v", err)
//	}
//	fmt.Println("Database for tests connected successfully!")
//}

//func InitDB() {
//	connStr := fmt.Sprintf(
//		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
//		os.Getenv("DB_HOST"),
//		os.Getenv("DB_PORT"),
//		os.Getenv("DB_USER"),
//		os.Getenv("DB_PASSWORD"),
//		os.Getenv("DB_NAME"),
//		os.Getenv("DB_SSLMODE"),
//	)
//
//	var err error
//	// Используем "=" вместо ":=" для обновления глобальной переменной DB
//	DB, err = sql.Open("postgres", connStr)
//	if err != nil {
//		log.Fatalf("Failed to connect to database: %v", err)
//	}
//
//	err = DB.Ping()
//	if err != nil {
//		log.Fatalf("Failed to ping database: %v", err)
//	}
//
//	fmt.Println("Database connected successfully!")
//}
