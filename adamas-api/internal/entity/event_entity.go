package entity

type Event struct {
	ID              int64
	Name            string
	Address         string
	Date            string
	Description     string
	Subscribers     []*CommonUserBasic
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
	ID               int        `json:"id"`
	Name             string     `json:"name"`
	QuantityProjects int        `json:"quantity_projects"`
	Projects         []*Project `json:"projects"`
}

type RepositoryInEvent struct {
	ID      string
	Project *Project
	Locale  string
}
