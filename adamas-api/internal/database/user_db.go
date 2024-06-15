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

func (ud *UserDB) CreateUser(name, email, password string) (*entity.CommonUserExtend, error) {
	user := entity.NewCommonUserExtend(name, email, password)
	_, err := ud.db.Exec("INSERT INTO COMMON_USER(name, email, password) VALUES( ?, ?, ?)",  user.USER.Name, user.USER.Email, user.USER.Password)
	if err != nil {
		return nil, err
	}
	err = ud.db.QueryRow("SELECT id FROM COMMON_USER WHERE email = ?", email).Scan(
		&user.USER.ID,    
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (ud * UserDB) LoginUser(email, password string) (*entity.CommonUserExtend, error) {
    var user entity.CommonUserExtend
    user.USER = &entity.User{}

    err := ud.db.QueryRow("SELECT id, name, email FROM COMMON_USER WHERE email = ? and password = ?", email, utils.EncriptKey(password)).Scan(
        &user.USER.ID, &user.USER.Name, &user.USER.Email,
    )
    if err != nil {
        return nil, err
    }
	user.USER.UserType = "common_user"

    return &user, nil
}
