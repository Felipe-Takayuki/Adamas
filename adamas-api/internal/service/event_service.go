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

func (es *EventService) CreateEvent(name, address, date, description string, institutionID int64) (*entity.Event, error) {
	event, err := es.EventDB.CreateEvent(name, address, date, description, institutionID)
	if err != nil {
		return nil, err 
	}
	return event,nil
}
func (es *EventService) GetEventByName(name string) ([]*entity.Event, error) {
	events, err := es.EventDB.GetEventByName(name)
	if err != nil {
		return nil, err 
	}
	return events, nil 
}
func (es *EventService) GetEvents() ([]*entity.Event, error) {
	events, err := es.EventDB.GetEvents()
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (es *EventService) GetSubscribersByEventID(eventID, ownerID int64) ([]*entity.CommonUserBasic, error) {
	subscribers, err := es.EventDB.GetSubscribersByEventID(eventID, ownerID)
	if err != nil {
		return nil, err 
	}
	return subscribers, nil
} 

func (es *EventService) EventRegistration(eventID, userID int64) ([]*entity.Event, error) {
	events, err := es.EventDB.EventRegistration(eventID,userID)
	if err != nil {
		return nil, err 
	}
	return events, nil
}

func (es *EventService) EventParticipation(eventID,userID, projectID int64) (*entity.Project, error){
	project, err := es.EventDB.EventParticipation(eventID,userID,projectID)
	if err != nil {
		return nil, err 
	}
	return project, nil
}

func (es *EventService) AddRoomInEvent(eventID int64, roomName string, quantityProjects int) ([]*entity.RoomEvent, error) {
	rooms, err := es.EventDB.AddRoomInEvent(eventID, roomName, quantityProjects)
	if err != nil {
		return nil, err 
	}
	return rooms, nil 
}