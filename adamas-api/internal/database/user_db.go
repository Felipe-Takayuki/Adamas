package database

import (
	"database/sql"

	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/entity"
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/utils"
)

type UserDB struct {
	db *sql.DB
}
func NewUserDB (db *sql.DB) *UserDB {
	return &UserDB{
		db: db,
	}
}

func (ud *UserDB) CreateUser(name, email, password string) (*entity.User, error) {
	user := entity.NewUser(name, email, password)
	_, err := ud.db.Exec("INSERT INTO COMMON_USER(name, email, password) VALUES( ?, ?, ?)",  user.Name, user.Email, user.Password)
	if err != nil {
		return nil, err
	}
	err = ud.db.QueryRow("SELECT id FROM COMMON_USER WHERE email = ?", email).Scan(
		&user.Id,    
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ud * UserDB) LoginUser(email, password string) (*entity.User, error) {
	var user entity.User
	err := ud.db.QueryRow("SELECT id, name, email FROM COMMON_USER WHERE email = ? and password = ?", email, utils.EncriptKey(password)).Scan(
		&user.Id, &user.Name, &user.Email,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}