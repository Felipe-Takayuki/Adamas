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

func (ud *UserDB) CreateUser(name, nickName, email, password string) (*entity.User, error) {
	user := entity.NewUser(name, nickName, email, password)
	result, err := ud.db.Exec(queries.CREATE_USER,  user.Name, user.NickName, user.Email, user.Password)
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

func (ud *UserDB) EditUser(name, nickName, description string, userID int64) (*entity.User, error) {
	
	if name != "" {
		_, err := ud.db.Exec(queries.UPDATE_USERNAME, name, userID)
		if err != nil {
			return nil, err 
		}
	}

	if nickName != "" {
		_, err := ud.db.Exec(queries.UPDATE_NICKNAME, nickName, userID)
		if err != nil {
			return nil, err 
		}
	}

	if description != "" {
		_, err := ud.db.Exec(queries.UPDATE_DESCRIPTION, description, userID)
		if err != nil {
			return nil, err 
		}
	}

	user := entity.User{Name: name, NickName: nickName, Description: description}
	return &user, nil 


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

func (ud *UserDB) GetUsers() ([]*entity.User, error) {
	rows, err  := ud.db.Query(queries.GET_USERS)
	if err != nil {
		return nil, err 
	}
	defer rows.Close()

	var users []*entity.User
	for rows.Next() {
		var user entity.User
		err := rows.Scan(&user.ID, &user.Name, &user.NickName, &user.Description)
		if err != nil {
			return nil, err 
		}
		users = append(users, &user)
	}
	return users, nil
}

func (ud *UserDB) GetUserByID(userID int64) (*entity.User, error) {
	var user entity.User 
	err := ud.db.QueryRow(queries.GET_USER, userID).Scan(&user.ID, &user.Name, &user.NickName, &user.Description)
	if err != nil {
		return nil, err 
	}
	return &user, nil  
}

func (ud *UserDB) GetUsersByName(name string) ([]*entity.User, error) {
	rows, err  := ud.db.Query(queries.GET_USERS_BY_NAME, name+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*entity.User
	for rows.Next() {
		var user entity.User
		err := rows.Scan(&user.ID, &user.Name, &user.NickName, &user.Description)
		if err != nil {
			return nil, err 
		}
		users = append(users, &user)
	}
	return users, nil	
}