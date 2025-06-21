package transaction_controller

import (
	"net/http"

	transaction_interfaces "github.com/agrawaltejas01/banks/internal/wallet/transaction/interfaces"
	transaction_model "github.com/agrawaltejas01/banks/internal/wallet/transaction/model"
	"github.com/gin-gonic/gin"
)

type TransactionController struct {
	service transaction_interfaces.TransactionService
}

func NewTransactionController(service transaction_interfaces.TransactionService) *TransactionController {
	return &TransactionController{service: service}
}

func (c *TransactionController) CreateTransaction(ctx *gin.Context) {
	var req struct {
		Amount   int    `json:"amount" binding:"required"`
		Type     string `json:"type" binding:"required"`
		WalletId string `json:"wallet_id" binding:"required"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tx, err := c.service.CreateTransaction(ctx, req.Amount, transaction_model.TransactionType(req.Type), req.WalletId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, tx)
}

func (c *TransactionController) GetTransactionById(ctx *gin.Context) {
	id := ctx.Param("id")
	tx, err := c.service.GetTransactionById(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}
	ctx.JSON(http.StatusOK, tx)
}
