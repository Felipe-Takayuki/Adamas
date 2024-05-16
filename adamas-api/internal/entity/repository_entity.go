package entity

type Bloc struct {
	SubTitle string
	Content  string
}

type Category struct {
	ID   int
	Name string
}

type Repository struct {
	ID                 int      `json:"id"`
	Title              string   `json:"title"`
	Description        string   `json:"description"`
	FirstOwnerUserID   int      `json:"user_id"`
	FirstOwnerUserName string   `json:"user_name"`
	OwnersID           []int    `json:"owners_id"`
	OwnersName         []string `json:"owners_name"`
	CategoriesID       []int	`json:"categories_id"`	
	CategoriesName     string	`json:"catetegories"`
	Blocs              []*Bloc	`json:"blocs"`
}

func NewRepository(title, description string, ownerID int) *Repository {
	return &Repository{
		Title:       title,
		Description: description,
		FirstOwnerUserID: ownerID,
	}
}
