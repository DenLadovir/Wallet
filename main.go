package main

import (
	"Wallet/database"
	"Wallet/handlers"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load("config.env")
	if err != nil {
		log.Fatalf("Error loading 'config.env' file: %v", err)
	}

	database.InitDB()

	r := mux.NewRouter()

	r.HandleFunc("/api/v1/wallets/{WALLET_UUID}/deposit", handlers.DepositHandler).Methods("POST")
	r.HandleFunc("/api/v1/wallets/{WALLET_UUID}/withdraw", handlers.WithdrawHandler).Methods("POST")
	r.HandleFunc("/api/v1/wallets/{WALLET_UUID}", handlers.GetWalletBalance).Methods("GET")
	r.HandleFunc("/api/v1/wallets", handlers.GetAllWalletsHandler).Methods("GET")

	fmt.Println("Server is running. Port - 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
