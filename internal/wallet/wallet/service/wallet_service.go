package wallet_service

import (
	"context"
	"errors"

	transaction_model "github.com/agrawaltejas01/banks/internal/wallet/transaction/model"
	user_interfaces "github.com/agrawaltejas01/banks/internal/wallet/user/interfaces"
	wallet_interfaces "github.com/agrawaltejas01/banks/internal/wallet/wallet/interfaces"
	wallet_model "github.com/agrawaltejas01/banks/internal/wallet/wallet/model"
)

type WalletService struct {
	repo        wallet_interfaces.WalletRepo
	userService user_interfaces.UserService
}

func NewWalletService(repo wallet_interfaces.WalletRepo, userService user_interfaces.UserService) *WalletService {
	return &WalletService{repo: repo, userService: userService}
}

func (s *WalletService) CreateWallet(userId string) (*wallet_model.Wallet, error) {

	user, err := s.userService.GetUserById(userId)
	if err != nil {
		return nil, err
	}

	wallet := wallet_model.NewWallet(user.Id)
	err = s.repo.Create(wallet)
	return wallet, err
}

func (s *WalletService) GetWalletById(ctx context.Context, id string) (*wallet_model.Wallet, error) {
	return s.repo.GetById(ctx, id)
}

func (s *WalletService) UpdateWalletBalance(ctx context.Context, id string, amount int,
	tType transaction_model.TransactionType) (*wallet_model.Wallet, error) {
	wallet, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	if tType == transaction_model.Debit {
		if wallet.Funds < amount {
			return nil, errors.New("insufficient balance")
		}
	}

	return wallet, s.repo.UpdateWalletBalance(ctx, id, amount, tType)
}
