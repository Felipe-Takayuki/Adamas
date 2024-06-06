package reqs

type CreateEventRequest struct {
	Name          string	`json:"name"`
	Address       string	`json:"address"`
	Date          string	`json:"date"`
	Description   string	`json:"description"`
}
