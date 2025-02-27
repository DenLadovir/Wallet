package handlers

import (
	"Wallet/database"
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// GetWalletBalance handles balance requests
func GetWalletBalance(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	walletID := vars["WALLET_UUID"]
	log.Printf("Balance request received for wallet ID: %s", walletID)

	if walletID == "" {
		log.Printf("Invalid wallet ID")
		http.Error(w, "Invalid wallet ID", http.StatusBadRequest)
		return
	}

	var balance float64
	err := database.DB.QueryRow(`SELECT balance FROM wallets WHERE id = $1`, walletID).Scan(&balance)
	if err == sql.ErrNoRows {
		log.Printf("Wallet not found: %s", walletID)
		http.Error(w, "Wallet not found", http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Database error: %v", err)
		http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("Balance for wallet ID %s: %.2f", walletID, balance)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"walletId": walletID,
		"balance":  balance,
	})
}
