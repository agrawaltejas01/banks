package server

import (
	user_controller "github.com/agrawaltejas01/banks/internal/wallet/user/controller"
	wallet_controller "github.com/agrawaltejas01/banks/internal/wallet/wallet/controller"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine, walletController *wallet_controller.WalletController, userController *user_controller.UserController) {
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
}
