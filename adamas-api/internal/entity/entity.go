package entity

import (
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/utils"
)

type UserTest struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	UserType string `json:"user_type"`
}
type InstitutionUserExtend struct {
	USER   *UserTest
	CNPJ   int      `json:"cnpj"`
	Events []*Event `json:"events"` 
}
func NewInstitutionUserExtend( name, email, password string ,cnpj int) *InstitutionUserExtend {
	user := &UserTest{Name: name, Email: email, Password: utils.EncriptKey(password), UserType: "institution"}
	return &InstitutionUserExtend{
		USER: user,
		CNPJ: cnpj,
	}

}
type CommonUserExtend struct {
	USER         *UserTest
	Repositories []*Repository `json:"repositories"`
}

func NewCommonUserExtend(name, email, password string ) *CommonUserExtend {
	return &CommonUserExtend{
		USER: &UserTest{Name: name, Email: email, Password: utils.EncriptKey(password), UserType: "common_user"},
	} 
}


