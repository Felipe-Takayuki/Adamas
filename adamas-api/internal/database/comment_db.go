package database

import (
	"fmt"

	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/entity"
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/utils/queries"
)

func (rdb *ProjectDB) SetComment(projectID, ownerID int64, comment string) error {
	_, err := rdb.db.Exec(queries.SET_COMMENT, ownerID, projectID, comment)
	if err != nil {
		return err
	}
	return nil
}
func (rdb *ProjectDB) DeleteComment(projectID, commentID int64) error {
	_, err := rdb.db.Exec(queries.DELETE_COMMENT, commentID, projectID)
	if err != nil {
		return err
	}
	return nil
}

func (rdb *ProjectDB) EditComment(text string, projectID, commentID, ownerID int64) (*entity.Comment, error) {
	if !rdb.isCommentOwner(projectID, ownerID) {
		return nil, fmt.Errorf("usuário não possui o repositório")
	}
	_, err := rdb.db.Exec(queries.UPDATE_COMMENT, text, projectID)
	if err != nil {
		return nil, err
	}
	comment, err := rdb.getCommentByID(commentID)
	if err != nil {
		return nil, err 
	}
	return comment, nil 
}

func (rdb *ProjectDB) getCommentByID(commentID int64) (*entity.Comment, error) {
	comment := &entity.Comment{}
	err := rdb.db.QueryRow(queries.GET_COMMENT_BY_ID, commentID).Scan(&comment.CommentID, &comment.UserID, &comment.UserName, &comment.Comment)
	if err != nil {
		return nil, err 
	}
	return comment, nil 
	
}
func (rdb *ProjectDB) getCommentsByRepoID(projectID int64) ([]*entity.Comment, error) {
	rows, err := rdb.db.Query(queries.GET_COMMENTS_BY_PROJECT, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []*entity.Comment
	for rows.Next() {
		var comment entity.Comment
		if err := rows.Scan(&comment.CommentID, &comment.UserID, &comment.UserName, &comment.Comment); err != nil {
			return nil, err
		}
		comments = append(comments, &comment)
	}
	return comments, nil
}

func (rdb *ProjectDB) deleteCommentsByProjectID(projectID int64) error {
	_, err := rdb.db.Exec("DELETE FROM COMMENT WHERE project_id = ?", projectID)
	if err != nil {
		return err
	}
	return nil
}

func (rdb *ProjectDB) isCommentOwner(userID, commentID int64) bool {
	var count int
	err := rdb.db.QueryRow(queries.CHECK_COMMENT_OWNER, userID, commentID).Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}
