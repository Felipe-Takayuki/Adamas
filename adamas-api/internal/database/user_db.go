package database

import (
	"database/sql"

	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/entity"
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
	rows, err := ud.db.Query("SELECT r.id, r.title, r.description FROM REPOSITORY r JOIN OWNERS_REPOSITORY o ON r.id = o.repository_id JOIN COMMON_USER u ON o.owner_id = u.idWHERE u.username = ?",username)
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
	_, err := ud.db.Exec("INSERT INTO common_user(id, name, email, password) VALUES(?, ?, ?, ?)", user.Id, user.Name, user.Email, user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}