package wallet_service

import (
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

func (s *WalletService) GetWalletById(id string) (*wallet_model.Wallet, error) {
	return s.repo.GetById(id)
}
