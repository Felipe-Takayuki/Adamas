package entity

type Bloc struct {
	SubTitle string
	Content string
}

type Category struct {
	ID string
	Name string
}


type Repository struct {
	ID string
	Title string
	OwnersID []string
	OwnersName []string
	Categories []*Category
	Blocs []*Bloc
}

