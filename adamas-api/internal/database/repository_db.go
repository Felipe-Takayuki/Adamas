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
func (rdb *RepoDB) getCategoriesByRepoID(repositoryID int64) ([]*entity.Category, error) {
	rows, err := rdb.db.Query(queries.GET_CATEGORIES_BY_REPO, repositoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var categories []*entity.Category
	for rows.Next() {
		var category entity.Category
		if err := rows.Scan(&category); err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}
	return categories, nil
}
func (rdb *RepoDB) SetCategory(categoryName string, repositoryID int64) error {
	_, err := rdb.db.Exec(queries.SET_CATEGORY, utils.Categories[categoryName], repositoryID)
	if err != nil {
		return err
	}
	return nil
}

func (rdb *RepoDB) SetComment(repositoryID, ownerID int64, comment string) error {
	_, err := rdb.db.Exec(queries.SET_COMMENT, ownerID, repositoryID, comment)
	if err != nil {
		return err
	}
	return nil
}
func (rdb *RepoDB) DeleteComment(repository_id, comment_id int64) error {
	_, err := rdb.db.Exec(queries.DELETE_COMMENT, comment_id, repository_id)
	if err != nil {
		return err 
	}
	return nil
}
func (rdb *RepoDB) getCommentsByRepoID(repositoryID int64) ([]*entity.Comment, error) {
	rows, err := rdb.db.Query(queries.GET_COMMENTS_BY_REPO, repositoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []*entity.Comment
	for rows.Next() {
		var comment entity.Comment
		if err := rows.Scan(&comment.CommentID, &comment.UserID, &comment.UserName, &comment.Comment); err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}
	return comments, nil

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

func (rdb *RepoDB) DeleteRepo(email, password string, repo_id int64) error {
	_, err := rdb.db.Exec(queries.DELETE_REPOSITORY, repo_id, email, utils.EncriptKey(password))
	if err != nil {
		return err
	}
	return nil
}