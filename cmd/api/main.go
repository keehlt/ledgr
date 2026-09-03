package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// store in slice
var transactions []Transaction = make([]Transaction, 0)

// creating struct
type Transaction struct {
	TransactionID string `json:"transaction_id"`
	UserID        string `json:"user_id"`
	Amount        int64  `json:"amt"`
	Currency      string `json:"currency"`
	IPAddress     string `json:"ip_address"`
	TimeStamp     int64  `json:"timestamp"`
}

// start the server
func main() {
	router := gin.Default()

	router.POST("/transactions", createTransactions)
	router.GET("/transactions", getTransactions)

	router.Run(":8080")
}

// retrieve json data and convert to struct
func createTransactions(c *gin.Context) {
	var transactionReq Transaction

	//unable to bind
	if err := c.ShouldBindJSON(&transactionReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	//check invalid transaction
	if !validateTransactions(transactionReq) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid transaction",
		})
		return
	}

	//call storeTransactions function to store the transaction in slice
	storeTransactions(transactionReq)

	//valid transaction binding
	c.JSON(http.StatusOK, gin.H{
		"message":     "data converted to struct successfully",
		"transaction": transactionReq,
	})
}

// store the transaction in slice
func storeTransactions(transactionReq Transaction) {
	transactions = append(transactions, transactionReq)
}

// retrieve the stored transactions
func getTransactions(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"transactions": transactions,
	})
}

// validate the transaction
func validateTransactions(transactionReq Transaction) bool {
	if transactionReq.Amount <= 0 {
		return false
	} else if transactionReq.UserID == "" {
		return false
	} else if transactionReq.TransactionID == "" {
		return false
	} else if transactionReq.Currency == "" {
		return false
	} else if transactionReq.IPAddress == "" {
		return false
	} else if transactionReq.TimeStamp <= 0 {
		return false
	} else {
		return true
	}
}
