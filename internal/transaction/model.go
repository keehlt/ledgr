package transaction

type Transaction struct {
	TransactionID string `json:"transaction_id"`
	UserID        string `json:"user_id"`
	Amount        int64  `json:"amt"`
	Currency      string `json:"currency"`
	IPAddress     string `json:"ip_address"`
	TimeStamp     int64  `json:"timestamp"`
}
