package dtos

type TransferDto struct {
	FromUserID string  `json:"from_user_id" binding:"required"`
	ToUserID   string  `json:"to_user_id" binding:"required"`
	Amount     float64 `json:"amount" binding:"required"`
}
