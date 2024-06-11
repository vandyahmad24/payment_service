package request

type CreateTransactionRequest struct {
	Amount        int64  `json:"amount" binding:"required"`
	Currency      string `json:"currency" binding:"required"`
	PaymentMethod string `json:"payment_method" binding:"required"`
	Description   string `json:"description"`
	UserId        string `json:"user_id"`
}
