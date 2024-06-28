package database

import (
	"database/sql"
	"fmt"

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
		repository.Categories, err = rdb.getCategoriesByRepoID(repository.ID)
		if err != nil {
			return nil, err
		}
		repository.Comments, err = rdb.getCommentsByRepoID(repository.ID)
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
		repository.Categories, err = rdb.getCategoriesByRepoID(repository.ID)
		if err != nil {
			return nil, err
		}
		repository.Comments, err = rdb.getCommentsByRepoID(repository.ID)
		if err != nil {
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

func (rdb *RepoDB) EditRepo(title, description, content string, repository_id int64) (*entity.RepositoryBasic, error) {

	if title != "" {
		_, err := rdb.db.Exec(queries.UPDATE_TITLE_REPOSITORY, title, repository_id)
		if err != nil {
			return nil, err
		}
	}

	if description != "" {
		_, err := rdb.db.Exec(queries.UPDATE_DESCRIPTION_REPOSITORY, description, repository_id)
		if err != nil {
			return nil, err
		}
	}

	if content != "" {
		_, err := rdb.db.Exec(queries.UPDATE_CONTENT_REPOSITORY, content, repository_id)
		if err != nil {
			return nil, err
		}
	}
	repository := entity.RepositoryBasic{ID: int(repository_id), Title: title, Description: description, Content: content}
	return &repository, nil
}

func (rdb *RepoDB) DeleteRepo(email, password string, repoID int64) error {
	userID, err := rdb.validateUser(email, password)
	if err != nil {
		return err
	}

	if !rdb.isRepositoryOwner(userID, repoID) {
		return fmt.Errorf("usuário não possui o repositório")
	}

	err = rdb.deleteOwnerRepository(userID, repoID)
	if err != nil {
		return err
	}

	err = rdb.deleteCommentsByRepoID(repoID)
	if err != nil {
		return err
	}

	err = rdb.deleteCategoriesByRepoID(repoID)
	if err != nil {
		return err
	}

	_, err = rdb.db.Exec(queries.DELETE_REPOSITORY, repoID)
	if err != nil {
		return err
	}

	return nil
}

func (rdb *RepoDB) validateUser(email, password string) (int64, error) {
	var userID int64
	err := rdb.db.QueryRow(queries.VALIDATE_USER, email, utils.EncriptKey(password)).Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func (rdb *RepoDB) isRepositoryOwner(userID, repoID int64) bool {
	var count int
	err := rdb.db.QueryRow(queries.CHECK_REPOSITORY_OWNER, userID, repoID).Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}

func (rdb *RepoDB) deleteOwnerRepository(userID, repoID int64) error {
	_, err := rdb.db.Exec(queries.DELETE_OWNER_REPO, userID, repoID)
	if err != nil {
		return err
	}
	return nil
}
