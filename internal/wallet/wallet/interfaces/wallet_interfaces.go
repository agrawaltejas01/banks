package wallet_interfaces

import (
	"context"

	transaction_model "github.com/agrawaltejas01/banks/internal/wallet/transaction/model"
	wallet_model "github.com/agrawaltejas01/banks/internal/wallet/wallet/model"
)

type WalletRepo interface {
	Create(wallet *wallet_model.Wallet) error
	GetById(ctx context.Context, id string) (*wallet_model.Wallet, error)
	UpdateWalletBalance(ctx context.Context, id string, amount int,
		tType transaction_model.TransactionType) error
}

type WalletService interface {
	CreateWallet(userId string) (*wallet_model.Wallet, error)
	GetWalletById(ctx context.Context, id string) (*wallet_model.Wallet, error)
	UpdateWalletBalance(ctx context.Context, id string, amount int, tType transaction_model.TransactionType) (*wallet_model.Wallet, error)
}

type WalletController interface {
	CreateWallet(ctx interface{})
	GetWalletById(ctx interface{})
}
