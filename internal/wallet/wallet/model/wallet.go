package wallet_model

import (
	"github.com/google/uuid"
)

type Wallet struct {
	Id          string `gorm:"primaryKey" json:"id"`
	UserId      string `json:"user_id"`
	Funds       int    `gorm:"default:0" json:"funds"`
	RewardFunds int    `gorm:"default:0" json:"reward_funds"`
}

func NewWallet(userId string) *Wallet {
	return &Wallet{
		Id:          uuid.New().String(),
		UserId:      userId,
		Funds:       0,
		RewardFunds: 0,
	}
}
