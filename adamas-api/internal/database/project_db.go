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

func (rdb *ProjectDB) GetProjectsByName(title string) ([]*entity.Project, error) {
	rows, err := rdb.db.Query(queries.GET_PROJECT_BY_NAME, title)
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
		project.Categories, err = rdb.getCategoriesByRepoID(project.ID)
		if err != nil {
			return nil, err
		}
		project.Comments, err = rdb.getCommentsByRepoID(project.ID)
		if err != nil {
			return nil, err
		}
		projects = append(projects, &project)
	}

	return projects, nil
}
func (rdb *ProjectDB) GetProjects() ([]*entity.Project, error) {
	rows, err := rdb.db.Query(queries.GET_PROJECTS)
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
		project.Categories, err = rdb.getCategoriesByRepoID(project.ID)
		if err != nil {
			return nil, err
		}
		project.Comments, err = rdb.getCommentsByRepoID(project.ID)
		if err != nil {
			return nil, err
		}
		projects = append(projects, &project)
	}
	return projects, nil
}

func (rdb *ProjectDB) CreateProject(title, description, content string, ownerID int) (*entity.Project, error) {
	project := entity.NewProject(title, description, content, ownerID)
	result, err := rdb.db.Exec(queries.CREATE_PROJECT, &project.Title, &project.Description, &project.Content)

	if err != nil {
		return nil, err
	}
	project.ID, err = result.LastInsertId()
	if err != nil {
		return nil, err
	}

	err = rdb.db.QueryRow(queries.GET_OWNER_NAME_BY_ID, project.FirstOwnerID).Scan(&project.FirstOwnerName)
	if err != nil {
		return nil, err
	}
	_, err = rdb.db.Exec(queries.SET_OWNER, &project.ID, &project.FirstOwnerID)
	if err != nil {
		return nil, err
	}
	var ownerNames []string
	project.OwnerNames = append(ownerNames, project.FirstOwnerName)
	return project, nil
}

func (rdb *ProjectDB) EditProject(title, description, content string, projectID int64) (*entity.ProjectBasic, error) {

	if title != "" {
		_, err := rdb.db.Exec(queries.UPDATE_TITLE_PROJECT, title, projectID)
		if err != nil {
			return nil, err
		}
	}

	if description != "" {
		_, err := rdb.db.Exec(queries.UPDATE_DESCRIPTION_PROJECT, description, projectID)
		if err != nil {
			return nil, err
		}
	}

	if content != "" {
		_, err := rdb.db.Exec(queries.UPDATE_CONTENT_PROJECT, content, projectID)
		if err != nil {
			return nil, err
		}
	}
	project := entity.ProjectBasic{Title: title, Description: description, Content: content}
	return &project, nil
}

func (rdb *ProjectDB) DeleteProject(email, password string, projectID int64) error {
	userID, err := rdb.validateUser(email, password)
	if err != nil {
		return err
	}

	if !rdb.isProjectOwner(userID, projectID) {
		return fmt.Errorf("usuário não possui o repositório")
	}

	err = rdb.deleteOwnerProject(userID, projectID)
	if err != nil {
		return err
	}

	err = rdb.deleteCommentsByProjectID(projectID)
	if err != nil {
		return err
	}

	err = rdb.deleteCategoriesByRepoID(projectID)
	if err != nil {
		return err
	}

	_, err = rdb.db.Exec(queries.DELETE_PROJECT, projectID)
	if err != nil {
		return err
	}

	return nil
}

func (rdb *ProjectDB) validateUser(email, password string) (int64, error) {
	var userID int64
	err := rdb.db.QueryRow(queries.VALIDATE_USER, email, utils.EncriptKey(password)).Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func (rdb *ProjectDB) isProjectOwner(userID, projectID int64) bool {
	var count int
	err := rdb.db.QueryRow(queries.CHECK_PROJECT_OWNER, userID, projectID).Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}

func (rdb *ProjectDB) deleteOwnerProject(userID, projectID int64) error {
	_, err := rdb.db.Exec(queries.DELETE_OWNER_PROJECT, userID, projectID)
	if err != nil {
		return err
	}
	return nil
}
