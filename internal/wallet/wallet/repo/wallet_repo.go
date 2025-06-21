package wallet_repo

import (
	"context"
	"fmt"

	"github.com/agrawaltejas01/banks/internal/database"
	transaction_model "github.com/agrawaltejas01/banks/internal/wallet/transaction/model"
	wallet_model "github.com/agrawaltejas01/banks/internal/wallet/wallet/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type WalletRepo struct {
	db *gorm.DB
}

func NewWalletRepo(db *gorm.DB) *WalletRepo {
	return &WalletRepo{db: db}
}

func (r *WalletRepo) Create(wallet *wallet_model.Wallet) error {
	return r.db.Create(wallet).Error
}

func (r *WalletRepo) GetById(ctx context.Context, id string) (*wallet_model.Wallet, error) {
	var wallet wallet_model.Wallet

	db := database.GetDbInstanceFromContextOrDB(ctx)

	err := db.
		Clauses(clause.Locking{Strength: clause.LockingStrengthUpdate}).
		First(&wallet, "id = ?", id).Error
	return &wallet, err
}

func (r *WalletRepo) UpdateWalletBalance(ctx context.Context, id string,
	amount int, tType transaction_model.TransactionType) error {

	db := database.GetDbInstanceFromContextOrDB(ctx)
	fmt.Print(db)

	if tType == transaction_model.Credit {
		return db.Model(&wallet_model.Wallet{}).
			Where("id = ?", id).
			Update("funds", gorm.Expr("funds + ?", amount)).
			Error
	}
	return db.Model(&wallet_model.Wallet{}).
		Where("id = ? AND funds >= ?", id, amount).
		Update("funds", gorm.Expr("funds - ?", amount)).
		Error

}
