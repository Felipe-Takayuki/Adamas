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

func (ud *UserDB) CreateUser(name, nickName, description, email, password string) (*entity.User, error) {
	user := entity.NewUser(name, nickName, description, email, password)
	result, err := ud.db.Exec(queries.CREATE_USER,  user.Name, user.NickName, user.Description,user.Email, user.Password)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	} else { 
		user.ID = id
	}
	return user, nil
}

func (ud * UserDB) LoginUser(email, password string) (*entity.User, error) {
    var user entity.User

    err := ud.db.QueryRow(queries.LOGIN_USER, email, utils.EncriptKey(password)).Scan(
        &user.ID, &user.Name, &user.Email,
    )
    if err != nil {
        return nil, err
    }
	user.UserType = "common_user"

    return &user, nil
}
