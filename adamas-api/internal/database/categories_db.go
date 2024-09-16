package database

import (
	"fmt"

	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/entity"
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/utils"
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/utils/queries"
)

func (pdb *ProjectDB) getCategoriesByRepoID(repositoryID int64) ([]*entity.Category, error) {
	rows, err := pdb.db.Query(queries.GET_CATEGORIES_BY_PROJECT, repositoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var categories []*entity.Category
	for rows.Next() {
		var category entity.Category
		if err := rows.Scan(&category.Name); err != nil {
			return nil, err
		}
		categories = append(categories, &category)
	}
	return categories, nil
}
func (pdb *ProjectDB) SetCategory(categoryName string, repositoryID int64) error {
	_, err := pdb.db.Exec(queries.SET_CATEGORY, utils.Categories[categoryName], repositoryID)
	if err != nil {
		return err
	}
	return nil
}

func (pdb *ProjectDB) deleteCategoriesByRepoID(repoID int64) error {
	_, err := pdb.db.Exec("DELETE FROM CATEGORY_PROJECT WHERE project_id = ?", repoID)
	if err != nil {
		return err
	}
	return nil
}

func (pdb *ProjectDB) DeleteCategory(projectID, ownerID int64, categoryName string) error {
	if !pdb.isProjectOwner(ownerID, projectID) {
		return fmt.Errorf("usuário não possui o repositório")
	}
	_, err := pdb.db.Exec("DELETE FROM CATEGORY_PROJECT WHERE project_id = ? AND category_id = ?", projectID, utils.Categories[categoryName])
	if err != nil {
		return err
	}
	return nil
}
