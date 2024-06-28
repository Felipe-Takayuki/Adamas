package database

import (
	"database/sql"

	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/entity"
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/utils"
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/utils/queries"
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
	result, err := ud.db.Exec(queries.CREATE_USER,  user.USER.Name, user.USER.Email, user.USER.Password)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	} else { 
		user.USER.ID = id
	}
	return user, nil
}

func (ud * UserDB) LoginUser(email, password string) (*entity.CommonUserExtend, error) {
    var user entity.CommonUserExtend
    user.USER = &entity.User{}

    err := ud.db.QueryRow(queries.LOGIN_USER, email, utils.EncriptKey(password)).Scan(
        &user.USER.ID, &user.USER.Name, &user.USER.Email,
    )
    if err != nil {
        return nil, err
    }
	user.USER.UserType = "common_user"

    return &user, nil
}
