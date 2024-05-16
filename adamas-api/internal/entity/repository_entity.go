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
	OwnerIDs           []int    `json:"owners_id"`
	OwnerNames         []string `json:"owners_name"`
	CategoriesID       []int	`json:"categories_id"`	
	CategoriesName     string	`json:"catetegories"`
	Blocs              []*Bloc	`json:"blocs"`
}
type RepositoryRequestFirst struct {
	Title              string   `json:"title"`
	Description        string   `json:"description"`
	FirstOwnerID 	   int 		`json:"owner_id"`
}
func NewRepository(title, description string, ownerID int) *Repository {
	var ownerIDs []int 
	return &Repository{
		Title:       title,
		Description: description,
		OwnerIDs: append(ownerIDs, ownerID),
	}
}
