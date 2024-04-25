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
	ID          int
	Title       string
	Description string
	OwnersID    []int
	OwnersName  []string
	CategoriesID []int 
	CategoriesName string
	Blocs       []*Bloc
}

func NewRepository(title, description string, ownerID int, categoriesID []int) *Repository {
	var ownersID []int
	return &Repository{
		Title:       title,
		Description: description,
		OwnersID:    append(ownersID, ownerID),
		CategoriesID:  categoriesID,
	}
}
