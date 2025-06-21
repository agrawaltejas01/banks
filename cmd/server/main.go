package main

import (
	"log"
	"os"

	"github.com/agrawaltejas01/banks/internal/database"
	"github.com/agrawaltejas01/banks/internal/server"
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

	walletRepo := wallet_repo.NewWalletRepo(db)
	walletService := wallet_service.NewWalletService(walletRepo)
	walletController := wallet_controller.NewWalletController(walletService)

	r := gin.Default()
	server.InitRoutes(r, walletController)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
