package reqs



type AddRoomRequest struct {
	Name             string `json:"name"`
	QuantityProjects int    `json:"quantity_projects"`
}


type ApproveProjectRequest struct {
	ProjectID int64 `json:"project_id"`
	RoomID    int64 `json:"room_id"`
}
