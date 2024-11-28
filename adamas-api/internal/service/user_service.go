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

func (us *UserService) CreateUser(name, nickName, email, password string) (*entity.User, error) {
	user, err := us.UserDB.CreateUser(name, nickName, email, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *UserService) EditUser(name, nickName, description string, userID int64)  (*entity.User, error) {
	user, err := us.UserDB.EditUser(name, nickName ,description, userID)
	if err != nil {
		return nil, err 
	}
	return user, nil 
}

func (us *UserService) LoginUser(email, password string) (*entity.User, error) {
	user, err := us.UserDB.LoginUser(email, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (us *UserService) GetUsers() ([]*entity.User, error) {
	users, err := us.UserDB.GetUsers()
	if err != nil {
		return nil, err 
	}
	return users, nil 
}

func (us *UserService) GetUserByID(userID int64) (*entity.User, error) {
	user, err := us.UserDB.GetUserByID(userID)
	if err != nil {
		return nil, err 
	}
	return user, nil 
}

func (us *UserService) GetUsersByName(name string) ([]*entity.User, error) {
	users, err := us.UserDB.GetUsersByName(name)
	if err != nil {
		return nil, err 
	}
	return users, nil 
}