package server

import (
	wallet_controller "github.com/agrawaltejas01/banks/internal/wallet/wallet/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, walletController *wallet_controller.WalletController) {
	walletGroup := r.Group("/wallet")
	{
		walletGroup.POST("/", walletController.CreateWallet)
		walletGroup.GET("/:id", walletController.GetWalletById)
	}
}
