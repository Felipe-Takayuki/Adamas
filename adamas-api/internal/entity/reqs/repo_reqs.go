package reqs

type SetCategoryRequest struct {
	RepositoryID int64 `json:"repository_id"`
	CategoryName   string `json:"category_name"`
}

type SetCommentRequest struct {
	RepositoryID int64 `json:"repository_id"`
	Comment string `json:"comment"` 
}

type RepositoryRequestFirst struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Content 	string 	`json:"content"`
}

type DeleteCommentRequest struct {
	CommentID int64 `json:"comment_id"`
	RepositoryID int64 `json:"repository_id"`
}