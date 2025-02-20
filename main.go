package main

import (
	"Wallet/config"
	"Wallet/database"
	"Wallet/handlers"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	// Загружаем переменные окружения из config.env
	cfg := config.LoadConfig()

	// Инициализируем базу данных
	database.InitDB(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	// Создаём маршрутизатор
	r := mux.NewRouter()

	// Регистрируем обработчики
	r.HandleFunc("/api/v1/wallets/{WALLET_UUID}/deposit", handlers.DepositHandler).Methods("POST")
	r.HandleFunc("/api/v1/wallets/{WALLET_UUID}/withdraw", handlers.WithdrawHandler).Methods("POST")
	r.HandleFunc("/api/v1/wallets/{WALLET_UUID}", handlers.GetWalletBalance).Methods("GET")
	r.HandleFunc("/api/v1/wallets", handlers.GetAllWalletsHandler).Methods("GET")

	// Запускаем сервер
	fmt.Printf("Server is running. Port - %s...\n", cfg.AppPort)
	log.Fatal(http.ListenAndServe(":"+cfg.AppPort, r))
}
