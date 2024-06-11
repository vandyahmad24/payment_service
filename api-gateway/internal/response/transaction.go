package response

type CreateTransactionResponse struct {
	TransactionId string `json:"transaction_id"`
	Status        string `json:"status"`
	CreatedAt     string `json:"created_at"`
}
