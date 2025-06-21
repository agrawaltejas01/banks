package transaction_repo

import (
	"context"

	"github.com/agrawaltejas01/banks/internal/database"
	transaction_model "github.com/agrawaltejas01/banks/internal/wallet/transaction/model"
	"gorm.io/gorm"
)

type TransactionRepo struct {
	db *gorm.DB
}

func NewTransactionRepo(db *gorm.DB) *TransactionRepo {
	return &TransactionRepo{db: db}
}

func (r *TransactionRepo) Create(ctx context.Context, tx *transaction_model.Transaction) error {
	db := database.GetDbInstanceFromContextOrDB(ctx)
	return db.Create(tx).Error
}

func (r *TransactionRepo) GetById(id string) (*transaction_model.Transaction, error) {
	var tx transaction_model.Transaction
	err := r.db.First(&tx, "id = ?", id).Error
	return &tx, err
}
