package user_interfaces

import user_model "github.com/agrawaltejas01/banks/internal/user/user/model"

type UserRepo interface {
	Create(user *user_model.User) error
	GetById(id string) (*user_model.User, error)
}

type UserService interface {
	CreateUser(name, email string) (*user_model.User, error)
	GetUserById(id string) (*user_model.User, error)
}

type UserController interface {
	CreateUser(ctx interface{})
	GetUserById(ctx interface{})
}
