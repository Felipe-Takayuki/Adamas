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

func (rs *RepositoryService) CreateRepo(title, description string, ownerID int) (*entity.Repository, error) {
	repo, err := rs.RepositoryDB.CreateRepo(title, description, ownerID)
	if err != nil {
		return nil, err
	}
	return repo, nil
}
