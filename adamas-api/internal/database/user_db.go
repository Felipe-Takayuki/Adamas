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

func (ud * UserDB) GetRepositoriesByUserName(username string) ([]*entity.Repository, error) {
	rows, err := ud.db.Query("SELECT r.id, r.title, r.description FROM REPOSITORY r JOIN OWNERS_REPOSITORY o ON r.id = o.repository_id JOIN COMMON_USER u ON o.owner_id = u.id WHERE u.name = ?",username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var repositories []*entity.Repository 
	for rows.Next() {
		var repository entity.Repository
		err := rows.Scan(&repository.ID, &repository.Title, &repository.Description) 
		if err != nil {
			return nil, err
		}
		repositories = append(repositories, &repository)
	}
	return repositories, nil
}

func (ud *UserDB) GetRepositoriesByName(name string)  ([]*entity.Repository, error) {
	rows, err := ud.db.Query("SELECT r.id, r.title, r.description FROM REPOSITORY r JOIN OWNERS_REPOSITORY o ON r.id = o.repository_id JOIN COMMON_USER u ON o.owner_id = u.id WHERE r.title = ?", name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var repositories []*entity.Repository 
	for rows.Next() {
		var repository entity.Repository
		err := rows.Scan(&repository.ID, &repository.Title, &repository.Description) 
		if err != nil {
			return nil, err
		}
		repositories = append(repositories, &repository)
	}
	return repositories, nil
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