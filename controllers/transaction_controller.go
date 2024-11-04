// controllers/auth_controller.go
package controllers

import (
	"net/http"

	"temp.project/transferbank/dtos"
	"temp.project/transferbank/services"

	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	TransactionService services.TransactionService
}

// controllers/transaction_controller.go
func (t *TransactionController) GetTransactionsReport(c *gin.Context) {
	userID := c.Param("user_id")
	report, err := t.TransactionService.GetTransactionsReport(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, report)
}
func (t *TransactionController) Transfer(c *gin.Context) {
	var transferDto dtos.TransferDto
	if err := c.ShouldBindJSON(&transferDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := t.TransactionService.QueueTransfer(&transferDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusAccepted, gin.H{"message": "Transfer in progress"})
}
func (t *TransactionController) Payment(c *gin.Context) {
	var paymentDto dtos.PaymentDto
	if err := c.ShouldBindJSON(&paymentDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := t.TransactionService.ProcessPayment(&paymentDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Payment processed"})
}
func (t *TransactionController) TopUp(c *gin.Context) {
	var topUpDto dtos.TopUpDto
	if err := c.ShouldBindJSON(&topUpDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := t.TransactionService.TopUp(&topUpDto)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Top up successful"})
}
