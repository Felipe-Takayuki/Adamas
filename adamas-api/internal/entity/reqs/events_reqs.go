package reqs

type CreateEventRequest struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	Description string `json:"description"`
}

type AddRoomRequest struct {
	Name             string `json:"name"`
	QuantityProjects int    `json:"quantity_projects"`
}

type AddPendingProjectRequest struct {
	ProjectID int64 `json:"project_id"`
}

type ApproveProjectRequest struct {
	ProjectID int64 `json:"project_id"`
	RoomID    int64 `json:"room_id"`
}
