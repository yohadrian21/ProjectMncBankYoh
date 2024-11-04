package repositories

import (
	"gorm.io/gorm"
	"temp.project/transferbank/dtos"
	"temp.project/transferbank/models"
)

// TransactionRepository defines the methods available for transaction persistence.
type TransactionRepository interface {
	CreateTransfer(transferDto *dtos.TransferDto) error
	CreateTransaction(transaction *models.Transaction) error
	GetTransactionsByUserID(userID string) ([]models.Transaction, error)
	// You can add other methods related to transactions here, e.g., GetTransfer, UpdateTransfer, etc.
}
type transactionRepository struct {
	db *gorm.DB
}

func (r *transactionRepository) CreateTransaction(transaction *models.Transaction) error {
	return r.db.Create(transaction).Error // Save the transaction to the database
}

// CreateTransfer implements TransactionRepository.
func (*transactionRepository) CreateTransfer(transferDto *dtos.TransferDto) error {
	panic("unimplemented")
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db}
}

func (r *transactionRepository) GetTransactionsByUserID(userID string) ([]models.Transaction, error) {
	var transactions []models.Transaction
	if err := r.db.Where("from_user_id = ? OR to_user_id = ?", userID, userID).Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
}
