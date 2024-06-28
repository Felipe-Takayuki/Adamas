package service

import (
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/database"
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/entity"
)

type RepositoryService struct {
	RepositoryDB database.RepoDB
}

func NewRepoService(repoDB database.RepoDB) *RepositoryService {
	return &RepositoryService{
		RepositoryDB: repoDB,
	}
}
func (rs *RepositoryService) GetRepositoriesByName(name string) ([]*entity.Repository, error) {
	repositories, err := rs.RepositoryDB.GetRepositoriesByName(name)
	if err != nil {
		return nil, err
	}
	return repositories, nil
}
func (rs *RepositoryService) GetRepositories()([]*entity.Repository, error) {
	repositories, err := rs.RepositoryDB.GetRepositories()
	if err != nil {
		return nil, err 
	}
	return repositories, nil
}
func (rs *RepositoryService) CreateRepo(title, description, content string, ownerID int) (*entity.Repository, error) {
	repo, err := rs.RepositoryDB.CreateRepo(title, description, content,ownerID)
	if err != nil {
		return nil, err
	}
	return repo, nil
}

func (rs *RepositoryService) EditRepo(title, description, content string, repository_id int64) (*entity.RepositoryBasic, error) {
	repo, err := rs.RepositoryDB.EditRepo(title, description, content, repository_id) 
	if err != nil {
		return nil, err
	}
	return repo, nil
}

func (rs *RepositoryService) DeleteRepo(email, password string, repositoryID int64) error {

	err := rs.RepositoryDB.DeleteRepo(email, password, repositoryID)
	if err != nil {
		return err
	}
	return nil
}

func (rs *RepositoryService) SetCategory(categoryName string, repoID int64) (error) {
	err := rs.RepositoryDB.SetCategory(categoryName, repoID)
	if err != nil {
		return err
	}
	return err 
}

