package service

func (rs *RepositoryService) SetComment(ownerID, repositoryID int64, comment string) (error) {
	err := rs.RepositoryDB.SetComment(repositoryID, ownerID, comment) 
	if err != nil {
		return err 
	}
	return nil
}

func (rs *RepositoryService) DeleteComment(comment_id, repository_id int64) (error) {
	err := rs.RepositoryDB.DeleteComment(repository_id, comment_id)
	if err != nil {
		return err 
	}
	return nil
}
