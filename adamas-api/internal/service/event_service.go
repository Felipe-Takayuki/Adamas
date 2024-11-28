package service

import (
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/database"
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/entity"
)

type EventService struct {
	EventDB *database.EventDB
}

func NewEventService(eventDB *database.EventDB) *EventService {
	return &EventService{
		EventDB: eventDB,
	}
}

func (es *EventService) CreateEvent(name, address, startDate, endDate, description string, institutionID int64) (*entity.Event, error) {
	event, err := es.EventDB.CreateEvent(name, address, startDate, endDate, description, institutionID)
	if err != nil {
		return nil, err
	}
	return event, nil
}
func (es *EventService) GetEventByName(name string) ([]*entity.Event, error) {
	events, err := es.EventDB.GetEventByName(name)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (es *EventService) GetEventByID(eventID int64) (*entity.Event, error) {
	event, err := es.EventDB.GetEventByID(eventID)
	if err != nil {
		return nil, err 
	}
	return event,nil 
}
func (es *EventService) GetEvents() ([]*entity.Event, error) {
	events, err := es.EventDB.GetEvents()
	if err != nil {
		return nil, err
	}
	return events, nil
}
func (es *EventService) GetEventByOwnerID(ownerID int64) ([]*entity.Event, error) {
	events, err := es.EventDB.GetEventByOwnerID(ownerID)
	if err != nil {
		return nil, err 
	}
	return events, nil
}

func (es *EventService) GetSubscribersByEventID(eventID, ownerID int64) ([]*entity.User, error) {
	subscribers, err := es.EventDB.GetSubscribersByEventID(eventID, ownerID)
	if err != nil {
		return nil, err
	}
	return subscribers, nil
}

func (es *EventService) GetPendingProjectsInEvent(eventID, ownerID int64) ([]*entity.Project, error) {
	pendingProjects, err := es.EventDB.GetPendingProjectsInEvent(eventID, ownerID) 
	if err != nil {
		return nil, err 
	}
	return pendingProjects, nil 
}


func (es *EventService) GetProjectsInEvent(eventID int64) ([]*entity.Project, error) {
	approvedProjects, err := es.EventDB.GetProjectsInEvent(eventID)
	if err != nil {
		return nil, err 
	}
	return approvedProjects, nil 
}


func (es *EventService) GetRoomsByEventID(eventID, owner_id int64) ([]*entity.RoomEvent, error) {
	rooms, err := es.EventDB.GetRoomsByEventID(eventID, owner_id)
	if err != nil {
		return nil, err 
	}
	return rooms, nil 
}
func (es *EventService) DeleteEvent(eventID int64, email, password string) error {
	err := es.EventDB.DeleteEvent(eventID,  email, password)
	if err != nil {
		return err 
	}
	return nil 
}

func (es *EventService) DeleteRoom(roomID, eventID int64) error {
	err := es.EventDB.DeleteRoom(roomID, eventID)
	if err != nil {
		return err 
	}
	return nil 
}

func (es *EventService) EventRegistration(eventID, userID int64) (*entity.Event, error) {
	event, err := es.EventDB.EventRegistration(eventID, userID)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (es *EventService) DeleteRegistrationInEvent(eventID, userID int64) (*entity.Event, error) {
	event, err := es.EventDB.DeleteRegistrationInEvent(eventID, userID)
	if err != nil {
		return nil, err 
	}
	return event, nil 
}

func (es *EventService) EventRequestParticipation(eventID, userID, projectID int64) (*entity.Project, error) {
	project, err := es.EventDB.EventRequestParticipation(eventID, userID, projectID)
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (es *EventService) DeleteParticipationInEvent(eventID, userID, projectID int64) (string, error) {
	participationWasDeleted, err := es.EventDB.DeleteParticipationInEvent(eventID, userID, projectID)
	if err != nil {
		return "", err 
	}
	return participationWasDeleted, nil 
}

func (es *EventService) ApproveParticipation(projectID, ownerID, eventID, roomID int64) ([]*entity.Project, error) {
	project, err := es.EventDB.ApproveParticipation(projectID, ownerID, eventID, roomID)
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (es *EventService) DisaApproveParticipation(projectID, eventID, ownerID int64) (string, error) {
	wasDisaaprove, err := es.EventDB.DisaApproveParticipation(projectID, eventID, ownerID)
	if err != nil {
		return "", err 
	}
	return wasDisaaprove, nil 
}

func (es *EventService) AddRoomInEvent(eventID, ownerID int64, roomName string, quantityProjects int) ([]*entity.RoomEvent, error) {
	rooms, err := es.EventDB.AddRoomInEvent(eventID, ownerID, roomName, quantityProjects)
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func (es *EventService) EditEvent(eventID, ownerID int64, name, address, startDate, endDate, description string) (*entity.Event, error) {
	event, err := es.EventDB.EditEvent(eventID, ownerID, name, address, startDate, endDate, description)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (es *EventService) EditRoom(roomID, eventID, quantityProjects, ownerID int64, roomName string) (*entity.RoomEvent, error) {
	room, err := es.EventDB.EditRoom(roomID, eventID, quantityProjects, ownerID, roomName) 
	if err != nil{
		return nil, err 
	}
	return room, nil 
}