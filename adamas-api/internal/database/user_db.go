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

func (ud * UserDB) GetRepositories( ) ([]*entity.Repository, error) {
	rows, err := ud.db.Query("SELECT id, title, description FROM repository")
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