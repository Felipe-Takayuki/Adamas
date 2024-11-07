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
func (ps *ProjectService) GetProjectsByName(name string) ([]*entity.Project, error) {
	projects, err := ps.ProjectDB.GetProjectsByName(name)
	if err != nil {
		return nil, err
	}
	return projects, nil
}
func (ps *ProjectService) GetProjectsByNameWithCategories(title string, categories []int64) ([]*entity.Project, error) {
	projects, err := ps.ProjectDB.GetProjectsByNameWithCategories(title, categories)
	if err != nil {
		return nil, err 
	}
	return projects, nil 
}
func (ps *ProjectService) GetProjectByID(projectID int64) (*entity.Project, error) {
	project, err := ps.ProjectDB.GetProjectByID(projectID)
	if err != nil {
		return nil, err 
	}
	return project, nil
}

func (ps *ProjectService) LikeProject(projectID, userID int64) ([]*entity.Like, error) {
	likes, err := ps.ProjectDB.LikeProject(projectID, userID)
	if err != nil {
		return nil, err 
	}
	return likes, nil 
}

func (ps *ProjectService) GetProjectsByCategories(categories  []int64) ([]*entity.Project, error) {
	projects, err := ps.ProjectDB.GetProjectsByCategorie(categories)
	if err != nil {
		return nil, err    
	}
	return projects, nil 
}
func (ps *ProjectService) GetProjects()([]*entity.Project, error) {
	repositories, err := ps.ProjectDB.GetProjects()
	if err != nil {
		return nil, err 
	}
	return repositories, nil
}
func (ps *ProjectService) CreateProject(title, description, content string, ownerID int) (*entity.Project, error) {
	repo, err := ps.ProjectDB.CreateProject(title, description, content,ownerID)
	if err != nil {
		return nil, err
	}
	return repo, nil
}

func (ps *ProjectService) EditProject(title, description, content string, projectID, ownerID int64) (*entity.Project, error) {
	repo, err := ps.ProjectDB.EditProject(title, description, content, projectID, ownerID) 
	if err != nil {
		return nil, err
	}
	return repo, nil
}

func (ps *ProjectService) DeleteProject(email, password string, projectID int64) error {

	err := ps.ProjectDB.DeleteProject(email, password, projectID)
	if err != nil {
		return err
	}
	return nil
}

func (ps *ProjectService) AddNewUserProject(projectID, userID, ownerID int64) ([]*entity.User, error) {
	useps, err := ps.ProjectDB.AddNewUserProject(projectID, userID, ownerID)
	if err != nil {
		return nil, err 
	}
	return useps, nil 
}
func (ps *ProjectService) SetCategory(categoryName string, projectID int64) (error) {
	err := ps.ProjectDB.SetCategory(categoryName, projectID)
	if err != nil {
		return err
	}
	return nil 
}
func (ps *ProjectService) DeleteCategory(projectID,ownerID int64, categoryID string) error {
	err := ps.ProjectDB.DeleteCategory(projectID, ownerID, categoryID)
	if err != nil {
		return err 
	}
	return nil 
}

func (ps *ProjectService) GetProjectsByUser(userID int64) ([]*entity.Project, error) {
	projects, err := ps.ProjectDB.GetProjectsByUser(userID)
	if err != nil {
		return nil, err 
	}
	return projects, nil
}