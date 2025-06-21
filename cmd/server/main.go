package main

import (
	"log"
	"os"

	"github.com/agrawaltejas01/banks/internal/database"
	"github.com/agrawaltejas01/banks/internal/server"
	transaction_controller "github.com/agrawaltejas01/banks/internal/wallet/transaction/controller"
	transaction_repo "github.com/agrawaltejas01/banks/internal/wallet/transaction/repo"
	transaction_service "github.com/agrawaltejas01/banks/internal/wallet/transaction/service"
	user_controller "github.com/agrawaltejas01/banks/internal/wallet/user/controller"
	user_repo "github.com/agrawaltejas01/banks/internal/wallet/user/repo"
	user_service "github.com/agrawaltejas01/banks/internal/wallet/user/service"
	wallet_controller "github.com/agrawaltejas01/banks/internal/wallet/wallet/controller"
	wallet_repo "github.com/agrawaltejas01/banks/internal/wallet/wallet/repo"
	wallet_service "github.com/agrawaltejas01/banks/internal/wallet/wallet/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}
	db := database.InitDB()
	database.AutoMigrate(db)

	userRepo := user_repo.NewUserRepo(db)
	userService := user_service.NewUserService(userRepo)
	userController := user_controller.NewUserController(userService)

	walletRepo := wallet_repo.NewWalletRepo(db)
	walletService := wallet_service.NewWalletService(walletRepo, userService)
	walletController := wallet_controller.NewWalletController(walletService)

	txnRepo := transaction_repo.NewTransactionRepo(db)
	txnService := transaction_service.NewTransactionService(txnRepo, walletService, userService)
	transactionController := transaction_controller.NewTransactionController(txnService)

	r := gin.Default()
	server.InitRoutes(r, walletController, userController, transactionController)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
