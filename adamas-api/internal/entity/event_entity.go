package entity

type Event struct {
	ID                    int64        `json:"id"`
	Name                  string       `json:"name"`
	Address               string       `json:"address"`
	StartDate             string       `json:"start_date"`
	EndDate               string       `json:"end_date"`
	Description           string       `json:"description"`
	Subscribers           []*User      `json:"subscribers,omitempty"`
	ProjectsParticipating []*Project   `json:"projects_participating,omitempty"`
	PendingProjects       []*Project   `json:"pending_projects,omitempty"`
	InstitutionID         int64        `json:"institution_id"`
	InstitutionName       string       `json:"institution_name"`
	Rooms                 []*RoomEvent `json:"rooms,omitempty"`
}

func NewEvent(name, address, startDate, endDate, description string, institutionID int64) *Event {
	return &Event{
		Name:          name,
		Address:       address,
		StartDate:     startDate,
		EndDate:       endDate,
		Description:   description,
		InstitutionID: institutionID,
	}
}

type RoomEvent struct {
	ID               int64      `json:"room_id"`
	Name             string     `json:"name"`
	QuantityProjects int        `json:"quantity_projects,omitempty"`
	Projects         []*Project `json:"projects,omitempty"`
}
