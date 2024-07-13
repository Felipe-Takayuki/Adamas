package reqs

type CreateEventRequest struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	Date        string `json:"date"`
	Description string `json:"description"`
}

type AddRoomRequest struct {
	Name             string `json:"name"`
	QuantityProjects int    `json:"quantity_projects"`
}
