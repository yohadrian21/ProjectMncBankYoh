// models/transaction.go
package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	FromUserID string
	ToUserID   string
	Amount     float64
	Status     string // e.g., Pending, Completed
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
	// Add other models
}
