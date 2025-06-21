package wallet_interfaces

import wallet_model "github.com/agrawaltejas01/banks/internal/wallet/wallet/model"

type WalletRepo interface {
	Create(wallet *wallet_model.Wallet) error
	GetById(id string) (*wallet_model.Wallet, error)
}

type WalletService interface {
	CreateWallet(userId string) (*wallet_model.Wallet, error)
	GetWalletById(id string) (*wallet_model.Wallet, error)
}

type WalletController interface {
	CreateWallet(ctx interface{})
	GetWalletById(ctx interface{})
}
