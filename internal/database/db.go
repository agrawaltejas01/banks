package database

import (
	"log"
	"os"

	transaction_model "github.com/agrawaltejas01/banks/internal/wallet/transaction/model"
	user_model "github.com/agrawaltejas01/banks/internal/wallet/user/model"
	wallet_model "github.com/agrawaltejas01/banks/internal/wallet/wallet/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		log.Fatal("MYSQL_DSN not set in environment")
	}
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	return db
}

func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(&wallet_model.Wallet{})
	if err != nil {
		log.Fatalf("AutoMigrate failed for wallet: %v", err)
	}

	err = db.AutoMigrate(&user_model.User{})
	if err != nil {
		log.Fatalf("AutoMigrate failed for user: %v", err)
	}

	err = db.AutoMigrate(&transaction_model.Transaction{})
	if err != nil {
		log.Fatalf("AutoMigrate failed for transaction: %v", err)
	}
}
