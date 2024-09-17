package entity

import "github.com/Felipe-Takayuki/Adamas/adamas-api/internal/utils"

type User struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	NickName    string `json:"nickname"`
	Description string `json:"description"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	UserType    string `json:"user_type"`
}
type Institution struct {
	ID          int64    `json:"id"`
	Name        string   `json:"name"`
	NickName    string   `json:"nickname"`
	Description string   `json:"description"`
	Email       string   `json:"email"`
	Password    string   `json:"password"`
	UserType    string   `json:"user_type"`
	CNPJ        string   `json:"cnpj"`
	Events      []*Event `json:"events"`
}

func NewInstitution(name, email, password string, cnpj string) *Institution {
	return &Institution{
		Name:     name,
		Email:    email,
		Password: utils.EncriptKey(password),
		CNPJ:     cnpj,
	}
}

func NewUser(name, nickName, description, email, password string) *User {
	return &User{
		Name:        name,
		NickName:    nickName,
		Description: description,
		Email:       email,
		Password:    utils.EncriptKey(password),
	}
}
