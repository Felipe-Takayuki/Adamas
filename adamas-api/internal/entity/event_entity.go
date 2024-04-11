package entity

type Event struct {
	ID string 
	Name string
	OwnerID string
	OwnerName string
	Repositories []*RepositoryInEvent
}

type RepositoryInEvent struct {
	ID string
	Repository *Repository
	Locale string
}

type Comment struct {
	UserID string
	UserName string
	Comment string
}