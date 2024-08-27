package reqs


type CategoryRequest struct {
	CategoryName string `json:"category_name"`
}

type SetCommentRequest struct {
	Comment string `json:"comment"`
}

type AddNewUserRequest struct {
	NewUserID int64 `json:"user_id"`
}

type EditCommentRequest struct {
	ID      int64  `json:"comment_id"`
	Comment string `json:"comment"`
}
type ProjectRequestFirst struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
}

type CommentID struct {
	CommentID int64 `json:"comment_id"`
}
