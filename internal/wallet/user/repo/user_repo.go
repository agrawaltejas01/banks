package user_repo

import (
	user_model "github.com/agrawaltejas01/banks/internal/wallet/user/model"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) Create(user *user_model.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepo) GetById(id string) (*user_model.User, error) {
	var user user_model.User
	err := r.db.First(&user, "id = ?", id).Error
	return &user, err
}
