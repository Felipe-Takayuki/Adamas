package service

import "github.com/Felipe-Takayuki/Adamas/adamas-api/internal/entity"

func (rs *ProjectService) SetComment(ownerID, repositoryID int64, comment string) error {
	err := rs.ProjectDB.SetComment(repositoryID, ownerID, comment)
	if err != nil {
		return err
	}
	return nil
}

func (rs *ProjectService) DeleteComment(comment_id, repository_id int64) error {
	err := rs.ProjectDB.DeleteComment(repository_id, comment_id)
	if err != nil {
		return err
	}
	return nil
}

func (rs *ProjectService) EditComment(text string, projectID, commentID, ownerID int64) (*entity.Comment, error) {
	comment, err := rs.ProjectDB.EditComment(text, projectID, commentID, ownerID)
	if err != nil {
		return nil, err 
	}
	return comment, nil 
}
