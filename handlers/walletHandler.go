package handlers

import (
	"Wallet/database"
	"Wallet/model"
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func DepositHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	walletID := vars["WALLET_UUID"]
	log.Printf("Deposit request received for wallet ID: %s", walletID)

	var req model.WalletRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.Amount <= 0 {
		log.Printf("Invalid request body: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
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

func WithdrawHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	walletID := vars["WALLET_UUID"]
	log.Printf("Withdraw request received for wallet ID: %s", walletID)

	var req model.WalletRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil || req.Amount <= 0 {
		log.Printf("Invalid request body: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	log.Printf("Withdraw amount: %.2f", req.Amount)

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

	if currentBalance < req.Amount {
		log.Printf("Insufficient funds for wallet ID %s. Current balance: %.2f, Withdraw amount: %.2f", walletID, currentBalance, req.Amount)
		http.Error(w, "Insufficient funds", http.StatusBadRequest)
		return
	}

	newBalance := currentBalance - req.Amount
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

	log.Printf("Withdraw successful for wallet ID %s. New balance: %.2f", walletID, newBalance)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"walletId": walletID,
		"balance":  newBalance,
	})
}

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

func GetAllWalletsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request received to fetch all wallets")

	rows, err := database.DB.Query(`SELECT id, balance, created_at FROM wallets`)
	if err != nil {
		log.Printf("Failed to fetch wallets: %v", err)
		http.Error(w, "Failed to fetch wallets: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var wallets []map[string]interface{}
	for rows.Next() {
		var id string
		var balance float64
		var createdAt string

		err := rows.Scan(&id, &balance, &createdAt)
		if err != nil {
			log.Printf("Failed to scan wallet: %v", err)
			http.Error(w, "Failed to scan wallet: "+err.Error(), http.StatusInternalServerError)
			return
		}

		wallet := map[string]interface{}{
			"id":         id,
			"balance":    balance,
			"created_at": createdAt,
		}
		wallets = append(wallets, wallet)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating over wallets: %v", err)
		http.Error(w, "Error iterating over wallets: "+err.Error(), http.StatusInternalServerError)
		return
	}

	log.Printf("Fetched %d wallets", len(wallets))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(wallets)
}
