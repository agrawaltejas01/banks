package user_model

import (
	"github.com/google/uuid"
)

type User struct {
	Id   string `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
}

func NewUser(name, email string) *User {
	return &User{
		Id:   uuid.New().String(),
		Name: name,
		Email: email,
	}
}
