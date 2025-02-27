package handlers

import (
	"Wallet/database"
	"Wallet/model"
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// DepositHandler handles deposit requests
func DepositHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	walletID := vars["WALLET_UUID"]
	log.Printf("Deposit request received for wallet ID: %s", walletID)
	if walletID == "" {
		log.Printf("Invalid wallet ID")
		http.Error(w, "Invalid wallet ID", http.StatusBadRequest)
		return
	}

	var req model.WalletRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.Amount <= 0 {
		log.Printf("Invalid request body: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	} else if req.Amount > 1000000 {
		log.Printf("Amount exceeds maximum limit: %.2f", req.Amount)
		http.Error(w, "Amount exceeds maximum limit", http.StatusBadRequest)
		return
	}
	log.Printf("Deposit amount: %.2f", req.Amount)

	tx, err := database.DB.Begin()
	if err != nil {
		log.Printf("Failed to start transaction: %v", err)
		http.Error(w, "Failed to start transaction: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer tx.Rollback()

	var currentBalance float64
	err = tx.QueryRow(`SELECT balance FROM wallets WHERE id = $1 FOR UPDATE`, walletID).Scan(&currentBalance)
	if err == sql.ErrNoRows {
		log.Printf("Wallet not found: %s", walletID)
		http.Error(w, "Wallet not found", http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Database error: %v", err)
		http.Error(w, "Database error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("Current balance for wallet ID %s: %.2f", walletID, currentBalance)

	newBalance := currentBalance + req.Amount
	_, err = tx.Exec(`UPDATE wallets SET balance = $1 WHERE id = $2`, newBalance, walletID)
	if err != nil {
		log.Printf("Failed to update balance: %v", err)
		http.Error(w, "Failed to update balance: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = tx.Commit()
	if err != nil {
		log.Printf("Failed to commit transaction: %v", err)
		http.Error(w, "Failed to commit transaction: "+err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("Deposit successful for wallet ID %s. New balance: %.2f", walletID, newBalance)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"walletId": walletID,
		"balance":  newBalance,
	})
}
