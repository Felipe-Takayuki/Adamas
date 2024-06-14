package database

import (
	"database/sql"

	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/entity"
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/utils"
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
	rows, err := rdb.db.Query("SELECT r.id, r.title, r.description, r.content, o.owner_id, u.name FROM REPOSITORY r JOIN OWNERS_REPOSITORY o ON r.id = o.repository_id JOIN COMMON_USER u ON o.owner_id = u.id WHERE r.title = ?", title)
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
	rows, err := rdb.db.Query("SELECT r.id, r.title, r.description,r.content, o.owner_id, u.name FROM REPOSITORY r JOIN OWNERS_REPOSITORY o ON r.id = o.repository_id JOIN COMMON_USER u ON o.owner_id = u.id")
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
	result, err := rdb.db.Exec("INSERT INTO REPOSITORY(title, description,content) VALUES (?,?,?)", &repo.Title, &repo.Description, &repo.Content)

	if err != nil {
		return nil, err
	}
	repo.ID, err = result.LastInsertId()
	if err != nil {
		return nil, err
	}

	err = rdb.db.QueryRow("SELECT name FROM COMMON_USER WHERE id = ?", repo.FirstOwnerID).Scan(&repo.FirstOwnerName)
	if err != nil {
		return nil, err
	}
	// _, err = rdb.db.Exec("INSERT INTO CATEGORY_REPO(repository_id) VALUES (?)", &repo.ID)
	// if err != nil {
	// 	return nil, err 
	// }

	_, err = rdb.db.Exec("INSERT INTO OWNERS_REPOSITORY(repository_id, owner_id) VALUES (?, ?)", &repo.ID, &repo.FirstOwnerID)
	if err != nil {
		return nil, err
	}
	var ownerNames []string
	repo.OwnerNames = append(ownerNames, repo.FirstOwnerName)

	return repo, nil
}

func (rdb *RepoDB) SetCategory(category_name string, repository_id int64) (error) {
	query := `
	UPDATE CATEGORY_REPO
	SET category_id = ?
	WHERE repository_id = ?`
	_, err := rdb.db.Exec(query,utils.Categories[category_name], repository_id)
	if err != nil {
		return err 
	} 
	return nil
}
