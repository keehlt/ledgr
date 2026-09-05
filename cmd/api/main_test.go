package main

import (
	"ledgr/internal/transaction"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

// test validateTransactions()
func TestValidateTransactions(t *testing.T) {

	tests := []struct {
		name        string
		transaction transaction.Transaction
		expected    bool
	}{
		{
			name: "valid transaction",
			transaction: transaction.Transaction{
				TransactionID: "tx001",
				UserID:        "user001",
				Amount:        5000,
				Currency:      "INR",
				IPAddress:     "127.0.0.1",
				TimeStamp:     1756940333,
			},
			expected: true,
		},
		{
			name: "negative amount",
			transaction: transaction.Transaction{
				TransactionID: "tx002",
				UserID:        "user002",
				Amount:        -5000,
				Currency:      "INR",
				IPAddress:     "127.0.0.2",
				TimeStamp:     1756940333,
			},
			expected: false,
		},
		{
			name: "empty transaction id",
			transaction: transaction.Transaction{
				TransactionID: "",
				UserID:        "user003",
				Amount:        5000,
				Currency:      "INR",
				IPAddress:     "127.0.0.3",
				TimeStamp:     1756940333,
			},
			expected: false,
		},
		{
			name: "empty user id",
			transaction: transaction.Transaction{
				TransactionID: "tx004",
				UserID:        "",
				Amount:        5000,
				Currency:      "INR",
				IPAddress:     "127.0.0.4",
				TimeStamp:     1756940333,
			},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := validateTransactions(test.transaction)

			if result != test.expected {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})

	}
}

// valid json
func TestCreateValidTransaction(t *testing.T) {
	router := gin.Default()
	router.POST("/transactions", createTransactions)

	request := httptest.NewRequest(
		http.MethodPost,
		"/transactions",
		strings.NewReader(`{"transaction_id":"tx-test","user_id":"user-test","amt":5000,"currency":"INR","ip_address":"127.0.0.1","timestamp":1756940333}`),
	)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)
	if response.Result().StatusCode != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, response.Result().StatusCode)
	}
}

// valid JSON, invalid transaction
func TestCreateInvalidTransaction(t *testing.T) {
	router := gin.Default()
	router.POST("/transactions", createTransactions)

	request := httptest.NewRequest(
		http.MethodPost,
		"/transactions",
		strings.NewReader(`{"transaction_id":"tx-test","user_id":"user-test","amt":-5000,"currency":"INR","ip_address":"127.0.0.1","timestamp":1756940333}`),
	)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	if response.Result().StatusCode != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, response.Result().StatusCode)
	}

}

// malformed JSON
func TestCreateInvalidJSON(t *testing.T) {
	router := gin.Default()
	router.POST("/transactions", createTransactions)

	request := httptest.NewRequest(
		http.MethodPost,
		"/transactions",
		strings.NewReader(`{"transaction_id":,"user_id":"user-test","amt":-5000,"currency":"INR","ip_address":"127.0.0.1","timestamp":1756940333}`),
	)
	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	if response.Result().StatusCode != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, response.Result().StatusCode)
	}
}

// test storage
func TestStoredTransaction(t *testing.T) {
	transaction := transaction.Transaction{
		TransactionID: "tx002",
		UserID:        "user002",
		Amount:        5000,
		Currency:      "INR",
		IPAddress:     "127.0.0.2",
		TimeStamp:     1756940333,
	}

	storeTransactions(transaction)
	y := len(transactions)
	st := transactions[y-1].TransactionID
	if st != transaction.TransactionID {
		t.Errorf("expected stored transaction ID %s, got %s", transaction.TransactionID, st)
	}
}
