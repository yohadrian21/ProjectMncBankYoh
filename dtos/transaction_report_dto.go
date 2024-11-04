// dtos/transaction_report_dto.go
package dtos

import "time"

type TransactionReportDto struct {
	TransactionID string    `json:"transaction_id"`
	Amount        float64   `json:"amount"`
	FromUserID    string    `json:"from_user_id"`
	ToUserID      string    `json:"to_user_id"`
	Timestamp     time.Time `json:"timestamp"` // Adjust the type based on your model
}
