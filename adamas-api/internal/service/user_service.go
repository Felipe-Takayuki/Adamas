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

func (us *UserService) GetRepositoriesByUserName(username string) ([]*entity.Repository, error) {
	repositories, err := us.UserDB.GetRepositoriesByUserName(username)
	if err != nil {
		return nil, err
	}
	return repositories, nil
}

func (us *UserService) CreateUser(name, email, password string) (*entity.User, error) {
	user, err := us.UserDB.CreateUser(name, email, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

