package service

import (
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/database"
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/entity"
)

type UserService struct {
	UserDB database.UserDB
}

func NewUserService(userDB database.UserDB) *UserService {
	return &UserService{
		UserDB: userDB,
	}
}

func (us *UserService) CreateUser(name, email, password string) (*entity.CommonUserExtend, error) {
	user, err := us.UserDB.CreateUser(name, email, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *UserService) LoginUser(email, password string) (*entity.CommonUserExtend, error) {
	user, err := us.UserDB.LoginUser(email, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}
