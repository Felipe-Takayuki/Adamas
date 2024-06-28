package entity

import (
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/utils"
)

type User struct {
	ID       int64    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	UserType string `json:"user_type"`
}
type InstitutionUserExtend struct {
	USER   *User
	CNPJ   int      `json:"cnpj"`
	Events []*Event `json:"events"` 
}
func NewInstitutionUserExtend( name, email, password string ,cnpj int) *InstitutionUserExtend {
	user := &User{Name: name, Email: email, Password: utils.EncriptKey(password), UserType: "institution_user"}
	return &InstitutionUserExtend{
		USER: user,
		CNPJ: cnpj,
	}

}
type CommonUserExtend struct {
	USER         *User
	Repositories []*Repository `json:"repositories"`
}

func NewCommonUserExtend(name, email, password string ) *CommonUserExtend {
	return &CommonUserExtend{
		USER: &User{Name: name, Email: email, Password: utils.EncriptKey(password), UserType: "common_user"},
	} 
}


