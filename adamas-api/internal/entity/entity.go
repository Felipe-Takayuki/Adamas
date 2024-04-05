package entity

import (
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/utils"
	"github.com/google/uuid"
)

type User struct {
	Id           string
	Name         string
	Email        string
	Password     string
	Repositories []*Repository
}

func NewUser(name string, email string, password string) *User {
	return &User{
		Id: uuid.NewString(),
		Name: name,
		Email: email,
		Password: utils.EncriptKey(password),
	}
}

type Institution struct {
	Name     string
	Email    string
	Password string
	CNPJ     string
	Events   []string
}
