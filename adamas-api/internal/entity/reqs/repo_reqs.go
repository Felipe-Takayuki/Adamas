package reqs

type SetCategoryRequest struct {
	CategoryName   string `json:"category_name"`
}

type SetCommentRequest struct {
	Comment string `json:"comment"` 
}

type RepositoryRequestFirst struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Content 	string 	`json:"content"`
}

type CommentID struct {
	CommentID int64 `json:"comment_id"`
}