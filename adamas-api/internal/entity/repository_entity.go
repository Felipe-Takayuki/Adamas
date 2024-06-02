package entity

type Category struct {
	ID   int
	Name string
}

type Repository struct {
	ID             int64         `json:"id"`
	Title          string      `json:"title"`
	Description    string      `json:"description"`
	Content 	   string 	   `json:"content"`
	FirstOwnerID   int         `json:"owner_id"`
	FirstOwnerName string      `json:"owner_name"`
	OwnerIDs       []int       `json:"owners_id"`
	OwnerNames     []string    `json:"owners_name"`
	Categories     []*Category `json:"categories"`
	Comments       []*Comment  `json:"comments"`
}
type ShowRepository struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
type RepositoryRequestFirst struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Content 	string 	`json:"content"`
}

func NewRepository(title, description, content string, ownerID int) *Repository {
	var ownerIDs []int
	return &Repository{
		Title:        title,
		Description:  description,
		Content: content,
		FirstOwnerID: ownerID,
		OwnerIDs:     append(ownerIDs, ownerID),
	}
}
