package database

import (
	"database/sql"

	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/entity"
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/utils"
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/utils/queries"
)

type RepoDB struct {
	db *sql.DB
}

func NewRepoDB(db *sql.DB) *RepoDB {
	return &RepoDB{
		db: db,
	}
}

func (rdb *RepoDB) GetRepositoriesByName(title string) ([]*entity.Repository, error) {
	rows, err := rdb.db.Query(queries.GET_REPOSITORY_BY_NAME, title)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var repositories []*entity.Repository
	for rows.Next() {
		var repository entity.Repository
		err := rows.Scan(&repository.ID, &repository.Title, &repository.Description, &repository.Content, &repository.FirstOwnerID, &repository.FirstOwnerName)
		if err != nil {
			return nil, err
		}
		repositories = append(repositories, &repository)
	}
	return repositories, nil
}
func (rdb *RepoDB) GetRepositories() ([]*entity.Repository, error) {
	rows, err := rdb.db.Query(queries.GET_REPOSITORIES)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var repositories []*entity.Repository
	for rows.Next() {
		var repository entity.Repository
		if err := rows.Scan(&repository.ID, &repository.Title, &repository.Description, &repository.Content, &repository.FirstOwnerID, &repository.FirstOwnerName); err != nil {
			return nil, err
		}
		repositories = append(repositories, &repository)
	}
	return repositories, nil
}

func (rdb *RepoDB) CreateRepo(title, description, content string, ownerID int) (*entity.Repository, error) {
	repo := entity.NewRepository(title, description, content, ownerID)
	result, err := rdb.db.Exec(queries.CREATE_REPOSITORY, &repo.Title, &repo.Description, &repo.Content)

	if err != nil {
		return nil, err
	}
	repo.ID, err = result.LastInsertId()
	if err != nil {
		return nil, err
	}

	err = rdb.db.QueryRow(queries.GET_OWNER_NAME_BY_ID, repo.FirstOwnerID).Scan(&repo.FirstOwnerName)
	if err != nil {
		return nil, err
	}
	_, err = rdb.db.Exec(queries.SET_OWNER, &repo.ID, &repo.FirstOwnerID)
	if err != nil {
		return nil, err
	}
	var ownerNames []string
	repo.OwnerNames = append(ownerNames, repo.FirstOwnerName)

	return repo, nil
}

func (rdb *RepoDB) SetCategory(category_name string, repository_id int64) (error) {
	_, err := rdb.db.Exec(queries.SET_CATEGORY,utils.Categories[category_name], repository_id)
	if err != nil {
		return err 
	} 
	return nil
}
