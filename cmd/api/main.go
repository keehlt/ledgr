package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// creating struct
type Transaction struct {
	TransactionID string `json:"transaction_id"`
	UserID        string `json:"user_id"`
	Amount        int64  `json:"amt"`
	Currency      string `json:"currency"`
	IPAddress     string `json:"ip"`
	TimeStamp     int64  `json:"timestamp"`
}

// start the server
func main() {
	router := gin.Default()

	router.POST("/transactions", getTransactions)

	router.Run(":8080")
}

// retrieve json data and convert to struct
func getTransactions(c *gin.Context) {
	var transactionReq Transaction

	if err := c.ShouldBindJSON(&transactionReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "data converted to struct successfully",
		"transaction": transactionReq,
	})
}
