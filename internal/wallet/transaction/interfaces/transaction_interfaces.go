package transaction_interfaces

import transaction_model "github.com/agrawaltejas01/banks/internal/wallet/transaction/model"

type TransactionRepo interface {
	Create(tx *transaction_model.Transaction) error
	GetById(id string) (*transaction_model.Transaction, error)
}

type TransactionService interface {
	CreateTransaction(amount int, tType transaction_model.TransactionType, walletId string) (*transaction_model.Transaction, error)
	GetTransactionById(id string) (*transaction_model.Transaction, error)
}

type TransactionController interface {
	CreateTransaction(ctx interface{})
	GetTransactionById(ctx interface{})
}
