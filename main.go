// main.go
package main

import (
	"github.com/gin-gonic/gin"
	"temp.project/transferbank/config"
	"temp.project/transferbank/controllers"
	"temp.project/transferbank/models"
	"temp.project/transferbank/repositories"
	"temp.project/transferbank/services"
)

func main() {
	router := gin.Default()

	// Initialize DB connection (choose one based on your database)
	db, err := config.ConnectPostgres("host=localhost user=gorm dbname=gorm port=9920 sslmode=disable")
	// or for MySQL
	// db, err := config.ConnectMySQL("gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("failed to connect database")
	}

	// Migrate models
	db.AutoMigrate(&models.User{}, &models.Transaction{})

	// Repositories
	userRepository := repositories.NewUserRepository(db)
	transactionRepo := repositories.NewTransactionRepository(db)

	// Services
	authService := services.NewAuthService(userRepository)
	transactionService := services.NewTransactionService(transactionRepo, userRepository)

	// Controllers
	authController := controllers.AuthController{AuthService: authService}
	transactionController := controllers.TransactionController{TransactionService: transactionService}

	// Routes
	router.POST("/register", authController.Register)
	router.POST("/login", authController.Login)
	router.POST("/topup", transactionController.TopUp)
	router.POST("/payment", transactionController.Payment)
	router.POST("/transfer", transactionController.Transfer)
	router.GET("/transactions", transactionController.GetTransactionsReport)
	// router.PUT("/profile", authController.updateProfile)
	// Add routes for login, top-up, payment, etc.

	router.Run(":8080")
}
