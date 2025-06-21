package transaction_service

import (
	"context"

	"github.com/agrawaltejas01/banks/internal/database"
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

func (s *TransactionService) CreateTransaction(ctx context.Context,
	amount int, tType transaction_model.TransactionType, walletId string) (*transaction_model.Transaction, error) {

	ctx = database.StartTransaction(ctx)

	// Update wallet balance
	_, err := s.walletService.UpdateWalletBalance(ctx, walletId, amount, tType)
	if err != nil {
		err = database.RollbackTransaction(ctx)
		if err != nil {
			return nil, err
		}
		return nil, err
	}

	// Within same transaction create transaction and update wallet balance
	tx := transaction_model.NewTransaction(amount, tType, walletId)
	err = s.repo.Create(ctx, tx)
	if err != nil {
		err = database.RollbackTransaction(ctx)
		if err != nil {
			return nil, err
		}
		return nil, err
	}

	err = database.CommitTransaction(ctx)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (s *TransactionService) GetTransactionById(id string) (*transaction_model.Transaction, error) {
	return s.repo.GetById(id)
}
