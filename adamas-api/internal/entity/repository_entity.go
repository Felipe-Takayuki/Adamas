package entity

type Category struct {
	ID   int
	Name string
}

type Repository struct {
	ID             int         `json:"id"`
	Title          string      `json:"title"`
	Description    string      `json:"description"`
	FirstOwnerID   int         `json:"owner_id"`
	FirstOwnerName string      `json:"owner_name"`
	OwnerIDs       []int       `json:"owners_id"`
	OwnerNames     []string    `json:"owners_name"`
	Categories     []*Category `json:"categories"`
}
type RepositoryRequestFirst struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func NewRepository(title, description string, ownerID int) *Repository {
	var ownerIDs []int
	return &Repository{
		Title:        title,
		Description:  description,
		FirstOwnerID: ownerID,
		OwnerIDs:     append(ownerIDs, ownerID),
	}
}
