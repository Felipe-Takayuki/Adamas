package entity

import (
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/utils"
)

type User struct {
	ID           int           `json:"id"`
	Name         string        `json:"name"`
	Email        string        `json:"email"`
	Password     string        `json:"password"`
	Repositories []*Repository `json:"repositories"`
}

type UserTest struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	UserType string `json:"user_type"`
}
func NewUserTest ( name, email, password, userType string) *UserTest {
	return &UserTest{ Name: name, Email: email, Password: password, UserType: userType}
}
type InstitutionUserExtend struct {
	USER   *UserTest
	CNPJ   int      `json:"cnpj"`
	Events []*Event `json:"events"`
}

func NewInstitutionUserExtend(cnpj int, name, email, password, userType string) *InstitutionUserExtend {
	user := &UserTest{Name: name, Email: email, Password: utils.EncriptKey(password), UserType: userType}
	return &InstitutionUserExtend{
		USER: user,
		CNPJ: cnpj,
	}
}
func NewUser(name string, email string, password string) *User {
	return &User{
		Name:     name,
		Email:    email,
		Password: utils.EncriptKey(password),
	}
}

type Institution struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	CNPJ     string   `json:"cnpj"`
	Events   []*Event `json:"events"`
}

type InstitutionCreateRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	CNPJ     string `json:"cnpj"`
}

func NewInstitution(name, email, password, cnpj string) *Institution {
	return &Institution{
		Name:     name,
		Email:    email,
		Password: password,
		CNPJ:     cnpj,
	}
}
