package service

func (rs *ProjectService) SetComment(ownerID, repositoryID int64, comment string) (error) {
	err := rs.ProjectDB.SetComment(repositoryID, ownerID, comment) 
	if err != nil {
		return err 
	}
	return nil
}

func (rs *ProjectService) DeleteComment(comment_id, repository_id int64) (error) {
	err := rs.ProjectDB.DeleteComment(repository_id, comment_id)
	if err != nil {
		return err 
	}
	return nil
}
