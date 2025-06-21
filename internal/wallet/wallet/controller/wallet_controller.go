package wallet_controller

import (
	"net/http"

	wallet_interfaces "github.com/agrawaltejas01/banks/internal/wallet/wallet/interfaces"
	"github.com/gin-gonic/gin"
)

type WalletController struct {
	service wallet_interfaces.WalletService
}

func NewWalletController(service wallet_interfaces.WalletService) *WalletController {
	return &WalletController{service: service}
}

func (c *WalletController) CreateWallet(ctx *gin.Context) {
	var req struct {
		UserId string `json:"user_id" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	wallet, err := c.service.CreateWallet(req.UserId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, wallet)
}

func (c *WalletController) GetWalletById(ctx *gin.Context) {
	id := ctx.Param("id")
	wallet, err := c.service.GetWalletById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Wallet not found"})
		return
	}
	ctx.JSON(http.StatusOK, wallet)
}
