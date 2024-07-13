package service

import (
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/database"
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/entity"
)

type ProjectService struct {
	ProjectDB database.ProjectDB
}

func NewProjectService(repoDB database.ProjectDB) *ProjectService {
	return &ProjectService{
		ProjectDB: repoDB,
	}
}
func (rs *ProjectService) GetProjectsByName(name string) ([]*entity.Project, error) {
	repositories, err := rs.ProjectDB.GetProjectsByName(name)
	if err != nil {
		return nil, err
	}
	return repositories, nil
}
func (rs *ProjectService) GetProjects()([]*entity.Project, error) {
	repositories, err := rs.ProjectDB.GetProjects()
	if err != nil {
		return nil, err 
	}
	return repositories, nil
}
func (rs *ProjectService) CreateProject(title, description, content string, ownerID int) (*entity.Project, error) {
	repo, err := rs.ProjectDB.CreateProject(title, description, content,ownerID)
	if err != nil {
		return nil, err
	}
	return repo, nil
}

func (rs *ProjectService) EditProject(title, description, content string, projectID int64) (*entity.ProjectBasic, error) {
	repo, err := rs.ProjectDB.EditProject(title, description, content, projectID) 
	if err != nil {
		return nil, err
	}
	return repo, nil
}

func (rs *ProjectService) DeleteProject(email, password string, projectID int64) error {

	err := rs.ProjectDB.DeleteProject(email, password, projectID)
	if err != nil {
		return err
	}
	return nil
}

func (rs *ProjectService) SetCategory(categoryName string, projectID int64) (error) {
	err := rs.ProjectDB.SetCategory(categoryName, projectID)
	if err != nil {
		return err
	}
	return err 
}

func (rs *ProjectService) GetProjectsByUser(userID int64) ([]*entity.Project, error) {
	projects, err := rs.ProjectDB.GetProjectsByUser(userID)
	if err != nil {
		return nil, err 
	}
	return projects, nil
}