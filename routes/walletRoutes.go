package routes

import (
	"Wallet/handlers"
	"github.com/gorilla/mux"
)

func SetupWalletRoutes(r *mux.Router) {
	r.HandleFunc("/api/v1/wallets/{WALLET_UUID}/deposit", handlers.DepositHandler).Methods("POST")
	r.HandleFunc("/api/v1/wallets/{WALLET_UUID}/withdraw", handlers.WithdrawHandler).Methods("POST")
	r.HandleFunc("/api/v1/wallets/{WALLET_UUID}", handlers.GetWalletBalance).Methods("GET")
	r.HandleFunc("/api/v1/wallets", handlers.GetAllWalletsHandler).Methods("GET")
}
