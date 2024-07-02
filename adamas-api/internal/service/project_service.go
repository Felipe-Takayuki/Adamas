package service

import (
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/database"
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/entity"
)

type ProjectService struct {
	RepositoryDB database.ProjectDB
}

func NewProjectService(repoDB database.ProjectDB) *ProjectService {
	return &ProjectService{
		RepositoryDB: repoDB,
	}
}
func (rs *ProjectService) GetProjectsByName(name string) ([]*entity.Project, error) {
	repositories, err := rs.RepositoryDB.GetProjectsByName(name)
	if err != nil {
		return nil, err
	}
	return repositories, nil
}
func (rs *ProjectService) GetProjects()([]*entity.Project, error) {
	repositories, err := rs.RepositoryDB.GetProjects()
	if err != nil {
		return nil, err 
	}
	return repositories, nil
}
func (rs *ProjectService) CreateProject(title, description, content string, ownerID int) (*entity.Project, error) {
	repo, err := rs.RepositoryDB.CreateProject(title, description, content,ownerID)
	if err != nil {
		return nil, err
	}
	return repo, nil
}

func (rs *ProjectService) EditProject(title, description, content string, projectID int64) (*entity.ProjectBasic, error) {
	repo, err := rs.RepositoryDB.EditProject(title, description, content, projectID) 
	if err != nil {
		return nil, err
	}
	return repo, nil
}

func (rs *ProjectService) DeleteProject(email, password string, projectID int64) error {

	err := rs.RepositoryDB.DeleteProject(email, password, projectID)
	if err != nil {
		return err
	}
	return nil
}

func (rs *ProjectService) SetCategory(categoryName string, projectID int64) (error) {
	err := rs.RepositoryDB.SetCategory(categoryName, projectID)
	if err != nil {
		return err
	}
	return err 
}

