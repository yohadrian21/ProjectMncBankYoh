package dtos

type PaymentDto struct {
	UserID  string  `json:"user_id" binding:"required"`
	Amount  float64 `json:"amount" binding:"required"`
	PayerID string  `json:"payer_id" binding:"required"`
	PayeeID string  `json:"payee_id" binding:"required"`
}
