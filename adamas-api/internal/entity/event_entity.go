package entity

type Event struct {
	ID              int64
	Name            string
	Address         string
	Date            string
	Description     string
	InstitutionID   int
	InstitutionName string
	Rooms           *[]RoomEvent
}

func NewEvent(name, address, date, description string, institutionID int) *Event {
	return &Event{
		Name:          name,
		Address:       address,
		Date:          date,
		Description:   description,
		InstitutionID: institutionID,
	}
}

type RoomEvent struct {
	ID                   int
	Name                 string
	QuantityRepositories int
	Repositories         *[9]Repository
}

type RepositoryInEvent struct {
	ID         string
	Repository *Repository
	Locale     string
}

type Comment struct {
	UserID   string
	UserName string
	Comment  string
}
