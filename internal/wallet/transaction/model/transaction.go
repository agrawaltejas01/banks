package transaction_model

import (
	"github.com/google/uuid"
)

type TransactionType string

const (
	Debit  TransactionType = "debit"
	Credit TransactionType = "credit"
)

type Transaction struct {
	Id       string          `gorm:"primaryKey" json:"id"`
	Amount   int             `json:"amount"`
	Type     TransactionType `json:"type"`
	WalletId string          `json:"wallet_id"`
}

func NewTransaction(amount int, tType TransactionType, walletId string) *Transaction {
	return &Transaction{
		Id:       uuid.New().String(),
		Amount:   amount,
		Type:     tType,
		WalletId: walletId,
	}
}
