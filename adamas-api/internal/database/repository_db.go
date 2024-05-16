package database

import (
	"database/sql"

	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/entity"
)

type RepoDB struct {
	db *sql.DB
}

func NewRepoDB(rdb *sql.DB) *RepoDB {
	return &RepoDB{
		db: rdb,
	}
}

func (rdb *RepoDB) CreateRepo(title, description string,ownerID int,) (*entity.Repository, error) {
	repo := entity.NewRepository(title, description, ownerID)
	_, err := rdb.db.Exec("INSERT INTO repository(title, description) VALUES (?,?)", &repo.Title, &repo.Description)
	if err != nil {
		return nil, err
	}
	err = rdb.db.QueryRow("SELECT id FROM repository WHERE title = ?", &repo.Title).Scan(&repo.ID)
	if err != nil {
		return nil, err
	}
	err = rdb.db.QueryRow("SELECT name FROM common_user WHERE id = ?", ownerID).Scan(&repo.FirstOwnerUserName)
	if err != nil {
		return nil, err
	}
	_, err = rdb.db.Exec("INSERT INTO OWNERS_REPOSITORY(repository_id, ownerID) VALUES (?, ?)", &repo.ID, ownerID)
	if err != nil {
		return nil, err 
	}
	return repo, nil 
}
