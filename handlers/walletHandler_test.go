package handlers

import (
	"Wallet/database"
	"Wallet/model"
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

// Mock database setup
func setupMockDB() {
	database.InitTestDB() // Инициализация базы данных

	// Создание таблицы и добавление тестовых данных
	database.DB.Exec(`CREATE TABLE IF NOT EXISTS wallets (
		id UUID PRIMARY KEY,
		balance FLOAT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`)

	database.DB.Exec(`INSERT INTO wallets (id, balance) VALUES 
		('123e4567-e89b-12d3-a456-426614174000', 1000.00),
		('987e6543-e21b-45d3-b789-123456789abc', 500.00)`)
}

func teardownMockDB() {
	// Очистка базы данных после тестов
	database.DB.Exec(`DROP TABLE IF EXISTS wallets`)
}

func waitForDB() {
	for {
		if database.DB != nil {
			err := database.DB.Ping()
			if err == nil {
				log.Println("Database is ready!")
				break
			}
		}
		log.Println("Waiting for database to be ready...")
		time.Sleep(1 * time.Second)
	}
}

func TestMain(m *testing.M) {
	database.InitTestDB()
	code := m.Run()
	if database.DB != nil {
		database.DB.Close()
	}
	os.Exit(code)
}

//func TestDepositHandler(t *testing.T) {
//	waitForDB() // Ждем, пока база данных будет готова
//
//	// Подготовка тестовых данных
//	setupMockDB()
//	defer teardownMockDB()
//
//	// Тестовый запрос
//	reqBody := model.WalletRequest{
//		Amount: 500.00,
//	}
//	body, _ := json.Marshal(reqBody)
//
//	req := httptest.NewRequest(http.MethodPost, "/api/v1/wallets/123e4567-e89b-12d3-a456-426614174000/deposit", bytes.NewReader(body))
//	w := httptest.NewRecorder()
//
//	DepositHandler(w, req)
//
//	// Проверка ответа
//	resp := w.Result()
//	if resp.StatusCode != http.StatusOK {
//		t.Errorf("Expected status OK, got %v", resp.StatusCode)
//	}
//}

// Test DepositHandler
func TestDepositHandler(t *testing.T) {
	waitForDB()
	setupMockDB()
	defer teardownMockDB()

	reqBody := model.WalletRequest{
		Amount: 500.00,
	}
	body, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/wallets/123e4567-e89b-12d3-a456-426614174000/deposit", bytes.NewReader(body))
	w := httptest.NewRecorder()

	DepositHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got %v", resp.StatusCode)
	}

	var respBody map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&respBody)

	if respBody["balance"] != 1500.00 {
		t.Errorf("Expected balance 1500.00, got %v", respBody["balance"])
	}
}

// Test WithdrawHandler
func TestWithdrawHandler(t *testing.T) {
	waitForDB()
	setupMockDB()
	defer teardownMockDB()

	reqBody := model.WalletRequest{
		Amount: 200.00,
	}
	body, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/wallets/123e4567-e89b-12d3-a456-426614174000/withdraw", bytes.NewReader(body))
	w := httptest.NewRecorder()

	WithdrawHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got %v", resp.StatusCode)
	}

	var respBody map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&respBody)

	if respBody["balance"] != 800.00 {
		t.Errorf("Expected balance 800.00, got %v", respBody["balance"])
	}
}

// Test GetWalletBalance
func TestGetWalletBalance(t *testing.T) {
	setupMockDB()
	defer teardownMockDB()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/wallets/123e4567-e89b-12d3-a456-426614174000/balance", nil)
	w := httptest.NewRecorder()

	GetWalletBalance(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got %v", resp.StatusCode)
	}

	var respBody map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&respBody)

	if respBody["balance"] != 1000.00 {
		t.Errorf("Expected balance 1000.00, got %v", respBody["balance"])
	}
}

// Test GetAllWalletsHandler
func TestGetAllWalletsHandler(t *testing.T) {
	setupMockDB()
	defer teardownMockDB()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/wallets", nil)
	w := httptest.NewRecorder()

	GetAllWalletsHandler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status OK, got %v", resp.StatusCode)
	}

	var respBody []map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&respBody)

	if len(respBody) != 2 {
		t.Errorf("Expected 2 wallets, got %v", len(respBody))
	}
}
