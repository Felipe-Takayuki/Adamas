package entity

import (
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/utils"
)

type User struct {
	Id           int
	Name         string
	Email        string
	Password     string
	Repositories []*Repository
}

func NewUser(name string, email string, password string) *User {
	return &User{
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
