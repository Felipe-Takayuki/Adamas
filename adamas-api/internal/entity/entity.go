package entity

import (
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/utils"
)

type User struct {
	Id           int           `json:"id"`
	Name         string        `json:"name"`
	Email        string        `json:"email"`
	Password     string        `json:"password"`
	Repositories []*Repository `json:"repositories"`
}

func NewUser(name string, email string, password string) *User {
	return &User{
		Name:     name,
		Email:    email,
		Password: utils.EncriptKey(password),
	}
}

type Institution struct {
	ID 		 int 	  `json:"id"`
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	CNPJ     string   `json:"cnpj"`
	Events   []*Event `json:"events"`
}

func NewInstitution(name, email, password, cnpj string) *Institution {
	return &Institution{
		Name:     name,
		Email:    email,
		Password: password,
		CNPJ:     cnpj,
	}
}
