package entity

type Category string

type Repository struct {
	ID             int64       `json:"id"`
	Title          string      `json:"title"`
	Description    string      `json:"description"`
	Content        string      `json:"content"`
	FirstOwnerID   int         `json:"owner_id"`
	FirstOwnerName string      `json:"owner_name"`
	OwnerIDs       []int       `json:"owners_id"`
	OwnerNames     []string    `json:"owners_name"`
	Categories     []*Category `json:"categories"`
	Comments       []*Comment  `json:"comments"`
}

type Comment struct {
	CommentID int64  `json:"comment_id"`
	UserID    int64  `json:"user_id"`
	UserName  string `json:"user_name"`
	Comment   string `json:"comment"`
}

type RepositoryBasic struct {
	ID          int    `json:"repository_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
}

func NewRepository(title, description, content string, ownerID int) *Repository {
	var ownerIDs []int
	return &Repository{
		Title:        title,
		Description:  description,
		Content:      content,
		FirstOwnerID: ownerID,
		OwnerIDs:     append(ownerIDs, ownerID),
	}
}
