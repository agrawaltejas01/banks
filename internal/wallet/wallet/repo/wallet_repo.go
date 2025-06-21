package wallet_repo

import (
	wallet_model "github.com/agrawaltejas01/banks/internal/wallet/wallet/model"
	"gorm.io/gorm"
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

func (r *WalletRepo) GetById(id string) (*wallet_model.Wallet, error) {
	var wallet wallet_model.Wallet
	err := r.db.First(&wallet, "id = ?", id).Error
	return &wallet, err
}

func (r *WalletRepo) UpdateWalletBalance(id string, funds int) error {
	return r.db.Model(&wallet_model.Wallet{}).
		Where("id = ?", id).
		UpdateColumn("funds", funds).
		Error
}
