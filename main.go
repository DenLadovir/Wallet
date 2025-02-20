package main

import (
	"Wallet/database"
	"Wallet/handlers"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	// Загружаем переменные окружения из config.env
	err := godotenv.Load("config.env")
	if err != nil {
		log.Printf("Warning: Could not load config.env file, using environment variables instead")
	}

	// Получаем параметры из переменных окружения
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432")
	dbUser := getEnv("POSTGRES_USER", "user")             // Изменено на POSTGRES_USER
	dbPassword := getEnv("POSTGRES_PASSWORD", "password") // Изменено на POSTGRES_PASSWORD
	dbName := getEnv("POSTGRES_DB", "wallet_db")          // Изменено на POSTGRES_DB
	appPort := getEnv("APP_PORT", "8080")

	// Инициализируем базу данных
	database.InitDB(dbHost, dbPort, dbUser, dbPassword, dbName)

	// Создаём маршрутизатор
	r := mux.NewRouter()

	// Регистрируем обработчики
	r.HandleFunc("/api/v1/wallets/{WALLET_UUID}/deposit", handlers.DepositHandler).Methods("POST")
	r.HandleFunc("/api/v1/wallets/{WALLET_UUID}/withdraw", handlers.WithdrawHandler).Methods("POST")
	r.HandleFunc("/api/v1/wallets/{WALLET_UUID}", handlers.GetWalletBalance).Methods("GET")
	r.HandleFunc("/api/v1/wallets", handlers.GetAllWalletsHandler).Methods("GET")

	// Запускаем сервер
	fmt.Printf("Server is running. Port - %s...\n", appPort)
	log.Fatal(http.ListenAndServe(":"+appPort, r))
}

// getEnv возвращает значение переменной окружения или значение по умолчанию
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
