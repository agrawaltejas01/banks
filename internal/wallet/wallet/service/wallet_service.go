package wallet_service

import (
	wallet_interfaces "github.com/agrawaltejas01/banks/internal/wallet/wallet/interfaces"
	wallet_model "github.com/agrawaltejas01/banks/internal/wallet/wallet/model"
)

type WalletService struct {
	repo wallet_interfaces.WalletRepo
}

func NewWalletService(repo wallet_interfaces.WalletRepo) *WalletService {
	return &WalletService{repo: repo}
}

func (s *WalletService) CreateWallet(userId string) (*wallet_model.Wallet, error) {
	wallet := wallet_model.NewWallet(userId)
	err := s.repo.Create(wallet)
	return wallet, err
}

func (s *WalletService) GetWalletById(id string) (*wallet_model.Wallet, error) {
	return s.repo.GetById(id)
}
