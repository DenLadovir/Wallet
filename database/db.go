package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"time"
)

var DB *sql.DB

func InitDB(host, port, user, password, dbname string) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
	}

	DB.SetMaxOpenConns(25)                 // Максимальное количество открытых соединений
	DB.SetMaxIdleConns(10)                 // Максимальное количество соединений в режиме ожидания
	DB.SetConnMaxLifetime(5 * time.Minute) // Максимальное время жизни соединения

	err = DB.Ping()
	if err != nil {
		log.Printf("Failed to ping database: %v", err)
	}

	fmt.Println("Database connected successfully with connection pool!")
}
