package database

import (
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/entity"
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/utils"
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/utils/queries"
)

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

func (rdb *RepoDB) deleteCategoriesByRepoID(repoID int64) error {
	_, err := rdb.db.Exec("DELETE FROM CATEGORY_REPO WHERE repository_id = ?", repoID)
	if err != nil {
		return err 
	}
	return nil 
}