package main

import (
	"testing"
)

func TestValidateTransactions(t *testing.T) {

	tests := []struct {
		name        string
		transaction Transaction
		expected    bool
	}{
		{
			name: "valid transaction",
			transaction: Transaction{
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
			transaction: Transaction{
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
			transaction: Transaction{
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
			transaction: Transaction{
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
