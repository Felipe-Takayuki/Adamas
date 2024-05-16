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
	_, err := rdb.db.Exec("INSERT INTO REPOSITORY(title, description) VALUES (?,?)", &repo.Title, &repo.Description)
	if err != nil {
		return nil, err
	}
	err = rdb.db.QueryRow("SELECT id FROM REPOSITORY WHERE title = ? AND description = ?", repo.Title, repo.Description).Scan(&repo.ID)
	if err != nil {
		return nil, err
	}
	
	err = rdb.db.QueryRow("SELECT name FROM COMMON_USER WHERE id = ?", repo.FirstOwnerID).Scan(&repo.FirstOwnerName)
	if err != nil {
		return nil, err
	}

	_, err = rdb.db.Exec("INSERT INTO OWNERS_REPOSITORY(repository_id, owner_id) VALUES (?, ?)", &repo.ID, &repo.FirstOwnerID)
	if err != nil {
		return nil, err 
	}

	return repo, nil 
}
