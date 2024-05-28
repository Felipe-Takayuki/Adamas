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

func (rdb * RepoDB) GetRepositoriesByName(title string) ([]*entity.Repository, error) {
	rows, err := rdb.db.Query("SELECT r.id, r.title, r.description, o.owner_id, u.name FROM REPOSITORY r JOIN OWNERS_REPOSITORY o ON r.id = o.repository_id JOIN COMMON_USER u ON o.owner_id = u.id WHERE r.title = ?",title)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var repositories []*entity.Repository
	for rows.Next() {
		var repository entity.Repository
		err := rows.Scan(&repository.ID, &repository.Title, &repository.Description, &repository.FirstOwnerID, &repository.FirstOwnerName) 
		if err != nil {
			return nil, err
		}
		repositories = append(repositories, &repository)
	}
	return repositories, nil
}
func (rdb *RepoDB) GetRepositories() ([]*entity.Repository, error) {
	rows, err := rdb.db.Query("SELECT r.id, r.title, r.description, o.owner_id, u.name FROM REPOSITORY r JOIN OWNERS_REPOSITORY o ON r.id = o.repository_id JOIN COMMON_USER u ON o.owner_id = u.id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var repositories []*entity.Repository
	for rows.Next() {
		var repository entity.Repository
		if err := rows.Scan(&repository.ID, &repository.Title, &repository.Description, &repository.FirstOwnerID, &repository.FirstOwnerName); err != nil{
			return nil, err
		}
		repositories = append(repositories, &repository)
	}	
	return repositories, nil
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
	var ownerNames[]string
	repo.OwnerNames = append(ownerNames, repo.FirstOwnerName)

	return repo, nil 
}

