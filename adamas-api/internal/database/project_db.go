package database

import (
	"database/sql"
	"fmt"

	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/entity"
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/utils"
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/utils/queries"
)

type ProjectDB struct {
	db *sql.DB
}

func NewProjectDB(db *sql.DB) *ProjectDB {
	return &ProjectDB{
		db: db,
	}
}

func (pdb *ProjectDB) GetProjectsByName(title string) ([]*entity.Project, error) {
	rows, err := pdb.db.Query(queries.GET_PROJECT_BY_NAME, title)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var projects []*entity.Project
	for rows.Next() {
		var project entity.Project
		err := rows.Scan(&project.ID, &project.Title, &project.Description, &project.Content, &project.FirstOwnerID, &project.FirstOwnerName)
		if err != nil {
			return nil, err
		}
		project.Categories, err = pdb.getCategoriesByRepoID(project.ID)
		if err != nil {
			return nil, err
		}
		project.Comments, err = pdb.getCommentsByRepoID(project.ID)
		if err != nil {
			return nil, err
		}
		projects = append(projects, &project)
	}

	return projects, nil
}
func (pdb *ProjectDB) GetProjects() ([]*entity.Project, error) {
	rows, err := pdb.db.Query(queries.GET_PROJECTS)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var projects []*entity.Project
	for rows.Next() {
		var project entity.Project
		if err := rows.Scan(&project.ID, &project.Title, &project.Description, &project.Content, &project.FirstOwnerID, &project.FirstOwnerName); err != nil {
			return nil, err
		}
		project.Categories, err = pdb.getCategoriesByRepoID(project.ID)
		if err != nil {
			return nil, err
		}
		project.Comments, err = pdb.getCommentsByRepoID(project.ID)
		if err != nil {
			return nil, err
		}
		project.Owners, err = pdb.getOwnersByProjectID(project.ID)
		if err != nil {
			return nil, err
		}
		
		projects = append(projects, &project)
	}
	return projects, nil
}



func (pdb *ProjectDB) GetProjectsByUser(userID int64) ([]*entity.Project, error) {
	rows, err := pdb.db.Query(queries.GET_PROJECTS_BY_USER, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var projects []*entity.Project
	for rows.Next() {
		var project entity.Project
		if err := rows.Scan(&project.ID, &project.Title, &project.Description, &project.Content, &project.FirstOwnerID, &project.FirstOwnerName); err != nil {
			return nil, err
		}
		project.Categories, err = pdb.getCategoriesByRepoID(project.ID)
		if err != nil {
			return nil, err
		}
		project.Comments, err = pdb.getCommentsByRepoID(project.ID)
		if err != nil {
			return nil, err
		}
		project.Owners, err = pdb.getOwnersByProjectID(project.ID)
		if err != nil {
			return nil, err
		}
		projects = append(projects, &project)
	}
	return projects, nil
}

func (pdb *ProjectDB) CreateProject(title, description, content string, ownerID int) (*entity.Project, error) {
	project := entity.NewProject(title, description, content, ownerID)
	result, err := pdb.db.Exec(queries.CREATE_PROJECT, &project.Title, &project.Description, &project.Content, &ownerID)

	if err != nil {
		return nil, err
	}
	project.ID, err = result.LastInsertId()
	if err != nil {
		return nil, err
	}

	err = pdb.db.QueryRow(queries.GET_OWNER_NAME_BY_ID, project.FirstOwnerID).Scan(&project.FirstOwnerName)
	if err != nil {
		return nil, err
	}
	_, err = pdb.db.Exec(queries.SET_OWNER, &project.ID, &project.FirstOwnerID)
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (pdb *ProjectDB) EditProject(title, description, content string, projectID, ownerID int64) (*entity.Project, error) {

	if !pdb.isProjectOwner(projectID, ownerID) {
		return nil, fmt.Errorf("usuário não possui o repositório")
	}
	if title != "" {
		_, err := pdb.db.Exec(queries.UPDATE_TITLE_PROJECT, title, projectID)
		if err != nil {
			return nil, err
		}
	}

	if description != "" {
		_, err := pdb.db.Exec(queries.UPDATE_DESCRIPTION_PROJECT, description, projectID)
		if err != nil {
			return nil, err
		}
	}

	if content != "" {
		_, err := pdb.db.Exec(queries.UPDATE_CONTENT_PROJECT, content, projectID)
		if err != nil {
			return nil, err
		}
	}
	project := entity.Project{Title: title, Description: description, Content: content}
	return &project, nil
}

func (pdb *ProjectDB) DeleteProject(email, password string, projectID int64) error {
	userID, err := pdb.validateUser(email, password)
	if err != nil {
		return err
	}

	if !pdb.isProjectOwner(userID, projectID) {
		return fmt.Errorf("usuário não possui o repositório")
	}

	err = pdb.deleteOwnerProject(userID, projectID)
	if err != nil {
		return err
	}

	err = pdb.deleteCommentsByProjectID(projectID)
	if err != nil {
		return err
	}

	err = pdb.deleteCategoriesByRepoID(projectID)
	if err != nil {
		return err
	}

	_, err = pdb.db.Exec(queries.DELETE_PROJECT, projectID)
	if err != nil {
		return err
	}

	return nil
}

func (pdb *ProjectDB) AddNewUserProject(projectID, userID, ownerID int64) ([]*entity.User, error) {
	if !pdb.isProjectOwner(ownerID, projectID) {
		return nil, fmt.Errorf("usuário não possui o repositório")
	}
	_, err := pdb.db.Exec("INSERT INTO OWNERS_PROJECT(project_id, owner_id) VALUES (?, ?)", projectID, userID)
	if err != nil {
		return nil, err
	}
	participants, err := pdb.getOwnersByProjectID(projectID)
	if err != nil {
		return nil, err
	}
	return participants, nil
}

func (pdb *ProjectDB) validateUser(email, password string) (int64, error) {
	var userID int64
	err := pdb.db.QueryRow(queries.VALIDATE_USER, email, utils.EncriptKey(password)).Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func (pdb *ProjectDB) isProjectOwner(userID, projectID int64) bool {
	var count int
	err := pdb.db.QueryRow(queries.CHECK_PROJECT_OWNER, userID, projectID).Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}

func (pdb *ProjectDB) getProjectByID(projectID int64) (*entity.Project, error) {
	project := &entity.Project{}
	err := pdb.db.QueryRow(queries.GET_PROJECT_BY_ID, projectID).Scan(&project.ID, &project.Title, &project.Description, &project.Content, &project.FirstOwnerID, &project.FirstOwnerName)
	if err != nil {
		return nil, err
	}
	categories, err := pdb.getCategoriesByRepoID(projectID)
	if err != nil {
		return nil, err
	}
	comments, err := pdb.getCommentsByRepoID(projectID)
	if err != nil {
		return nil, err
	}
	project.Comments = comments
	project.Categories = categories

	return project, nil
}

func (pdb *ProjectDB) deleteOwnerProject(userID, projectID int64) error {
	_, err := pdb.db.Exec(queries.DELETE_OWNER_PROJECT, userID, projectID)
	if err != nil {
		return err
	}
	return nil
}

func (pdb *ProjectDB) getOwnersByProjectID(projectID int64) ([]*entity.User, error) {
	rows, err := pdb.db.Query(queries.GET_OWNERS_BY_PROJECT, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var owners []*entity.User
	for rows.Next() {
		var owner entity.User
		err = rows.Scan(&owner.ID, &owner.Name)
		if err != nil {
			return nil, err
		}
		owners = append(owners, &owner)
	}
	return owners, nil
}
