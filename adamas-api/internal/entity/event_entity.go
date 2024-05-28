package entity

type Event struct {
	ID int 
	Name string
	OwnerID string
	OwnerName string
	Repositories []*RepositoryInEvent
}

type RoomEvent struct{
	ID int
	Name string
	QuantityRepositories int
	Repositories *[9]Repository
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

