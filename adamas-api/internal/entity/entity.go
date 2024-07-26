package entity

import (
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/utils"
)

type User struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	NickName    string `json:"nickname"`
	Description string `json:"description"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	UserType    string `json:"user_type"`
}
type InstitutionUserExtend struct {
	USER   *User
	CNPJ   string   `json:"cnpj"`
	Events []*Event `json:"events"`
}

func NewInstitutionUserExtend(name, email, password string, cnpj string) *InstitutionUserExtend {
	user := &User{Name: name, Email: email, Password: utils.EncriptKey(password), UserType: "institution_user"}
	return &InstitutionUserExtend{
		USER: user,
		CNPJ: cnpj,
	}

}

type CommonUserExtend struct {
	USER         *User
	Repositories []*Project `json:"projects"`
}
type CommonUserBasic struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func NewCommonUserExtend(name, nickName, description, email, password string) *CommonUserExtend {
	return &CommonUserExtend{
		USER: &User{Name: name, NickName: nickName, Description: description, Email: email, Password: utils.EncriptKey(password), UserType: "common_user"},
	}
}
