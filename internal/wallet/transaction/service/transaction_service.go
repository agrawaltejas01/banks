package transaction_service

import (
	transaction_interfaces "github.com/agrawaltejas01/banks/internal/wallet/transaction/interfaces"
	transaction_model "github.com/agrawaltejas01/banks/internal/wallet/transaction/model"
	user_service "github.com/agrawaltejas01/banks/internal/wallet/user/service"
	wallet_service "github.com/agrawaltejas01/banks/internal/wallet/wallet/service"
)

type TransactionService struct {
	repo          transaction_interfaces.TransactionRepo
	walletService *wallet_service.WalletService
	userService   *user_service.UserService
}

func NewTransactionService(repo transaction_interfaces.TransactionRepo, walletService *wallet_service.WalletService, userService *user_service.UserService) *TransactionService {
	return &TransactionService{repo: repo, walletService: walletService, userService: userService}
}

func (s *TransactionService) CreateTransaction(amount int, tType transaction_model.TransactionType, walletId string) (*transaction_model.Transaction, error) {

	// Update wallet balance
	_, err := s.walletService.UpdateWalletBalance(walletId, amount, tType)
	if err != nil {
		return nil, err
	}

	// Within same transaction create transaction and update wallet balance
	tx := transaction_model.NewTransaction(amount, tType, walletId)
	err = s.repo.Create(tx)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (s *TransactionService) GetTransactionById(id string) (*transaction_model.Transaction, error) {
	return s.repo.GetById(id)
}
