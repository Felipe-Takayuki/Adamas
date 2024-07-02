package entity

type Event struct {
	ID              int64
	Name            string
	Address         string
	Date            string
	Description     string
	InstitutionID   int64
	InstitutionName string
	Rooms           []*RoomEvent
}

func NewEvent(name, address, date, description string, institutionID int64) *Event {
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
	Projects         []*Project
}

type RepositoryInEvent struct {
	ID         string
	Project *Project
	Locale     string
}


