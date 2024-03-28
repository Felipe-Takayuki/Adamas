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

func (us *UserService) GetRepositories() ([]*entity.Repository, error) {
	repositories, err := us.UserDB.GetRepositories()
	if err != nil {
		return nil, err
	}
	return repositories, nil
}