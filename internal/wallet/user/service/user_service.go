package user_service

import (
	user_interfaces "github.com/agrawaltejas01/banks/internal/wallet/user/interfaces"
	user_model "github.com/agrawaltejas01/banks/internal/wallet/user/model"
)

type UserService struct {
	repo user_interfaces.UserRepo
}

func NewUserService(repo user_interfaces.UserRepo) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(name, email string) (*user_model.User, error) {
	user := user_model.NewUser(name, email)
	err := s.repo.Create(user)
	return user, err
}

func (s *UserService) GetUserById(id string) (*user_model.User, error) {
	return s.repo.GetById(id)
}
