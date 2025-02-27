package main

import (
	"Wallet/config"
	"Wallet/database"
	"Wallet/routes"
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	// Загружаем переменные окружения из config.env
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Printf("Ошибка при загрузке конфигурации: %v", err)
	}

	// Инициализируем базу данных
	database.InitDB(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	// Создаём маршрутизатор
	r := mux.NewRouter()

	// Регистрируем маршруты
	routes.SetupWalletRoutes(r)

	// Запускаем сервер
	fmt.Printf("Server is running. Port - %s...\n", cfg.AppPort)
	log.Fatal(http.ListenAndServe(":"+cfg.AppPort, r))
}
