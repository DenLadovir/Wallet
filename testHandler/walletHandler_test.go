package testHandler

import (
	"Wallet/database"
	"Wallet/handlers"
	"Wallet/model"
	"Wallet/testDB"
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
	testDB.InitTestDB() // Инициализация базы данных

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
	// Инициализируем тестовую базу данных
	testDB.InitTestDB() // Убедитесь, что InitTestDB использует значения по умолчанию
	defer func() {
		if err := database.DB.Close(); err != nil {
			log.Fatalf("Failed to close test database connection: %v", err)
		}
	}()

	// Очищаем и заполняем тестовую базу данных
	testDB.ClearTestDB()
	testDB.SeedTestDB([]map[string]interface{}{
		{"id": "123e4567-e89b-12d3-a456-426614174000", "balance": 1000.00},
		{"id": "223e4567-e89b-12d3-a456-426614174001", "balance": 500.00},
	})

	// Запускаем тесты
	code := m.Run()

	// Завершаем выполнение
	os.Exit(code)
}

func TestDatabaseConnection(t *testing.T) {
	err := database.DB.Ping()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
}

func TestClearTestDB(t *testing.T) {
	testDB.ClearTestDB()
}

func TestSeedTestDB(t *testing.T) {
	testDB.SeedTestDB([]map[string]interface{}{
		{"id": "123e4567-e89b-12d3-a456-426614174000", "balance": 1000.00},
		{"id": "223e4567-e89b-12d3-a456-426614174001", "balance": 500.00},
	})
}

// Test DepositHandler
func TestDepositHandler(t *testing.T) {
	waitForDB()
	setupMockDB()
	defer teardownMockDB()

	tests := []struct {
		name            string
		walletID        string
		requestBody     model.WalletRequest
		expectedStatus  int
		expectedBalance float64
	}{
		{
			name:            "Valid deposit",
			walletID:        "123e4567-e89b-12d3-a456-426614174000",
			requestBody:     model.WalletRequest{Amount: 500.00},
			expectedStatus:  http.StatusOK,
			expectedBalance: 1500.00,
		},
		{
			name:            "Invalid wallet ID",
			walletID:        "",
			requestBody:     model.WalletRequest{Amount: 500.00},
			expectedStatus:  http.StatusBadRequest,
			expectedBalance: 0,
		},
		{
			name:            "Negative deposit amount",
			walletID:        "123e4567-e89b-12d3-a456-426614174000",
			requestBody:     model.WalletRequest{Amount: -100.00},
			expectedStatus:  http.StatusBadRequest,
			expectedBalance: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.requestBody)
			req := httptest.NewRequest(http.MethodPost, "/api/v1/wallets/"+tt.walletID+"/deposit", bytes.NewReader(body))
			w := httptest.NewRecorder()

			handlers.DepositHandler(w, req)

			resp := w.Result()
			if resp.StatusCode != tt.expectedStatus {
				t.Errorf("Expected status %v, got %v", tt.expectedStatus, resp.StatusCode)
			}

			if tt.expectedStatus == http.StatusOK {
				var respBody map[string]interface{}
				json.NewDecoder(resp.Body).Decode(&respBody)

				if respBody["balance"] != tt.expectedBalance {
					t.Errorf("Expected balance %v, got %v", tt.expectedBalance, respBody["balance"])
				}
			}
		})
	}
}

// Test WithdrawHandler
func TestWithdrawHandler(t *testing.T) {
	waitForDB()
	setupMockDB()
	defer teardownMockDB()

	tests := []struct {
		name            string
		walletID        string
		requestBody     model.WalletRequest
		expectedStatus  int
		expectedBalance float64
	}{
		{
			name:            "Valid withdraw",
			walletID:        "123e4567-e89b-12d3-a456-426614174000",
			requestBody:     model.WalletRequest{Amount: 200.00},
			expectedStatus:  http.StatusOK,
			expectedBalance: 800.00,
		},
		{
			name:            "Insufficient funds",
			walletID:        "123e4567-e89b-12d3-a456-426614174000",
			requestBody:     model.WalletRequest{Amount: 2000.00},
			expectedStatus:  http.StatusBadRequest,
			expectedBalance: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.requestBody)
			req := httptest.NewRequest(http.MethodPost, "/api/v1/wallets/"+tt.walletID+"/withdraw", bytes.NewReader(body))
			w := httptest.NewRecorder()

			handlers.WithdrawHandler(w, req)

			resp := w.Result()
			if resp.StatusCode != tt.expectedStatus {
				t.Errorf("Expected status %v, got %v", tt.expectedStatus, resp.StatusCode)
			}

			if tt.expectedStatus == http.StatusOK {
				var respBody map[string]interface{}
				json.NewDecoder(resp.Body).Decode(&respBody)

				if respBody["balance"] != tt.expectedBalance {
					t.Errorf("Expected balance %v, got %v", tt.expectedBalance, respBody["balance"])
				}
			}
		})
	}
}
