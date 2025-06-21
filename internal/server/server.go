package server

import (
	transaction_controller "github.com/agrawaltejas01/banks/internal/wallet/transaction/controller"
	user_controller "github.com/agrawaltejas01/banks/internal/wallet/user/controller"
	wallet_controller "github.com/agrawaltejas01/banks/internal/wallet/wallet/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, walletController *wallet_controller.WalletController, userController *user_controller.UserController, transactionController *transaction_controller.TransactionController) {
	walletGroup := r.Group("/wallet")
	{
		walletGroup.POST("/", walletController.CreateWallet)
		walletGroup.GET("/:id", walletController.GetWalletById)
	}
	userGroup := r.Group("/user")
	{
		userGroup.POST("/", userController.CreateUser)
		userGroup.GET("/:id", userController.GetUserById)
	}
	transactionGroup := r.Group("/transaction")
	{
		transactionGroup.POST("/", transactionController.CreateTransaction)
		transactionGroup.GET("/:id", transactionController.GetTransactionById)
	}
}
