package handlers

import (
	"Wallet/database"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// GetAllWalletsHandler fetches all wallets
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
