package entity

import "github.com/Felipe-Takayuki/Adamas/adamas-api/internal/utils"

type User struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	NickName    string `json:"nickname,omitempty"`
	Description string `json:"description,omitempty"`
	Email       string `json:"email,omitempty"`
	Password    string `json:"password,omitempty"`
	UserType    string `json:"user_type,omitempty"`
}
type Institution struct {
	ID       int64    `json:"id"`
	Name     string   `json:"name"`
	Email    string   `json:"email,omitempty"`
	Password string   `json:"password,omitempty"`
	UserType string   `json:"user_type,omitempty"`
	CNPJ     string   `json:"cnpj,omitempty"`
	Events   []*Event `json:"events,omitempty"`
}

func NewInstitution(name, email, password string, cnpj string) *Institution {
	return &Institution{
		Name:     name,
		Email:    email,
		UserType: "institution_user",
		Password: utils.EncriptKey(password),
		CNPJ:     cnpj,
	}

}

func NewUser(name, email, password string) *User {
	return &User{
		Name:     name,
		Email:    email,
		UserType: "common_user",
		Password: utils.EncriptKey(password),
	}
}
