package database

import (
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/entity"
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/utils/queries"
)

func (rdb *RepoDB) SetComment(repositoryID, ownerID int64, comment string) error {
	_, err := rdb.db.Exec(queries.SET_COMMENT, ownerID, repositoryID, comment)
	if err != nil {
		return err
	}
	return nil
}
func (rdb *RepoDB) DeleteComment(repositoryID, commentID int64) error {
	_, err := rdb.db.Exec(queries.DELETE_COMMENT, commentID, repositoryID)
	if err != nil {
		return err
	}
	return nil
}

func (rdb *RepoDB) getCommentsByRepoID(repositoryID int64) ([]*entity.Comment, error) {
	rows, err := rdb.db.Query(queries.GET_COMMENTS_BY_REPO, repositoryID)
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
