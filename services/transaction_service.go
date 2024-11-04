// services/transaction_service.go
package services

import (
	"errors"
	"log"

	"temp.project/transferbank/dtos"

	"temp.project/transferbank/models"

	"temp.project/transferbank/repositories"
)

type TransactionService interface {
	// Here you could include a database connection or other dependencies
	// db *gorm.DB // Example if using GORM
	QueueTransfer(transferDto *dtos.TransferDto) error
	ProcessTransfer(transferDto *dtos.TransferDto)
	GetTransactionsReport(userID string) ([]dtos.TransactionReportDto, error)
	ProcessPayment(paymentDto *dtos.PaymentDto) error
	TopUp(topUpDto *dtos.TopUpDto) error
}

type transactionService struct {
	TransactionRepository repositories.TransactionRepository
	UserRepository        repositories.UserRepository
}

func NewTransactionService(transactionRepo repositories.TransactionRepository, userRepo repositories.UserRepository) TransactionService {
	return &transactionService{TransactionRepository: transactionRepo, UserRepository: userRepo}
}

func (t *transactionService) QueueTransfer(transferDto *dtos.TransferDto) error {
	go t.ProcessTransfer(transferDto)
	return nil
}

func (t *transactionService) ProcessTransfer(transferDto *dtos.TransferDto) {
	// Save the transfer using the repository
	if err := t.TransactionRepository.CreateTransfer(transferDto); err != nil {
		log.Printf("Failed to create transfer: %v", err)
		return
	}

	// Perform the transfer, update the database, etc.
	log.Println("Transfer processed successfully")
}

func (t *transactionService) GetTransactionsReport(userID string) ([]dtos.TransactionReportDto, error) {
	// Retrieve transactions from the repository
	transactions, err := t.TransactionRepository.GetTransactionsByUserID(userID)
	if err != nil {
		return nil, err
	}

	// Convert transactions to report DTOs
	var report []dtos.TransactionReportDto
	for _, transaction := range transactions {
		reportDto := dtos.TransactionReportDto{
			// TransactionID: transaction.ID,
			Amount:     transaction.Amount,
			FromUserID: transaction.FromUserID,
			ToUserID:   transaction.ToUserID,
			// Timestamp:     transaction.Timestamp, // Assuming you have this field
		}
		report = append(report, reportDto)
	}

	return report, nil
}

func (t *transactionService) ProcessPayment(paymentDto *dtos.PaymentDto) error {
	// Fetch payer user
	var payer models.User
	if err := t.UserRepository.FindByID(paymentDto.PayerID, &payer); err != nil {
		return errors.New("payer not found")
	}

	// Check if payer has sufficient balance
	if payer.Balance < paymentDto.Amount {
		return errors.New("insufficient balance")
	}

	// Fetch payee user
	var payee models.User
	if err := t.UserRepository.FindByID(paymentDto.PayeeID, &payee); err != nil {
		return errors.New("payee not found")
	}

	// Deduct amount from payer's balance and add to payee's balance
	payer.Balance -= paymentDto.Amount
	payee.Balance += paymentDto.Amount

	// Update balances in the database
	if err := t.UserRepository.Update(&payer); err != nil {
		return err
	}
	if err := t.UserRepository.Update(&payee); err != nil {
		return err
	}

	// Create transaction record
	transaction := models.Transaction{
		Amount:     paymentDto.Amount,
		FromUserID: paymentDto.PayerID,
		ToUserID:   paymentDto.PayeeID,
	}

	if err := t.TransactionRepository.CreateTransaction(&transaction); err != nil {
		return err
	}

	return nil
}

func (t *transactionService) TopUp(topUpDto *dtos.TopUpDto) error {
	// Fetch the user to top up
	var user models.User
	if err := t.UserRepository.FindByID(topUpDto.UserID, &user); err != nil {
		return errors.New("user not found")
	}

	// Update the user's balance
	user.Balance += topUpDto.Amount // Add the top-up amount to the user's balance

	// Update the user's balance in the database
	if err := t.UserRepository.Update(&user); err != nil {
		return err
	}

	// Create a transaction record for the top-up (optional)
	transaction := models.Transaction{
		Amount:     topUpDto.Amount,
		FromUserID: topUpDto.UserID, // Assuming the user is topping up their own account
		ToUserID:   topUpDto.UserID, // Same as above
	}

	if err := t.TransactionRepository.CreateTransaction(&transaction); err != nil {
		return err
	}

	return nil
}
