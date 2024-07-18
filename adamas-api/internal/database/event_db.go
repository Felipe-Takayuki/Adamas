package database

import (
	"database/sql"
	"fmt"

	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/entity"
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/utils/queries"
)

type EventDB struct {
	db *sql.DB
}

func NewEventDB(db *sql.DB) *EventDB {
	return &EventDB{
		db: db,
	}
}

func (edb *EventDB) CreateEvent(name, address, date, description string, institutionID int64) (*entity.Event, error) {
	event := entity.NewEvent(name, address, date, description, institutionID)
	result, err := edb.db.Exec(queries.CREATE_EVENT, &event.Name, &event.Address, &event.Date, &event.Description)
	if err != nil {
		return nil, err
	}
	event.ID, err = result.LastInsertId()
	if err != nil {
		return nil, err
	}
	_, err = edb.db.Exec(queries.SET_OWNER_EVENT, &event.ID, &event.InstitutionID)
	if err != nil {
		return nil, err
	}
	return event, nil
}
func (edb *EventDB) AddRoomInEvent(eventID int64, roomName string, quantityProjects int) ([]*entity.RoomEvent, error) {
	_, err := edb.db.Exec(queries.AddRoomInEvent, eventID, roomName, quantityProjects)
	if err != nil {
		return nil, err
	}
	rooms, err := edb.getRoomsByEventID(eventID)
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func (edb *EventDB) GetEventByName(name string) ([]*entity.Event, error) {
	rows, err := edb.db.Query(queries.GET_EVENT_BY_NAME, name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []*entity.Event
	for rows.Next() {
		var event entity.Event
		err = rows.Scan(&event.ID, &event.Name, &event.Address, &event.Date, &event.Description, &event.InstitutionID, &event.InstitutionName)
		if err != nil {
			return nil, err
		}
		rooms, err := edb.getRoomsByEventID(event.ID)
		if err != nil {
			return nil, err
		}
		event.Rooms = rooms
		events = append(events, &event)
	}
	return events, err
}
func (edb *EventDB) GetEvents() ([]*entity.Event, error) {
	rows, err := edb.db.Query(queries.GET_EVENTS)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []*entity.Event
	for rows.Next() {
		var event entity.Event
		err = rows.Scan(&event.ID, &event.Name, &event.Address, &event.Date, &event.Description, &event.InstitutionID, &event.InstitutionName)
		rooms, err := edb.getRoomsByEventID(event.ID)
		if err != nil {
			return nil, err
		}
		subscribers, err := edb.getSubscribersByEventID(event.ID)
		if err != nil {
			return nil, err
		}
		event.Subscribers = subscribers
		event.Rooms = rooms
		events = append(events, &event)
	}
	return events, err
}

func (edb *EventDB) EventRegistration(eventID, userID int64) ([]*entity.Event, error) {
	_, err := edb.db.Exec("INSERT INTO SUBSCRIBERS_EVENT(event_id, user_id) VALUES (?, ?)", eventID, userID)
	if err != nil {
		return nil, err
	}
	event, err := edb.getEventByID(eventID)
	if err != nil {
		return nil, err
	}
	return event, nil
}
func (edb *EventDB) EventParticipation(eventID, userID, projectID int64) (*entity.Project, error) {
	pdb := NewProjectDB(edb.db)
	if !pdb.isProjectOwner(userID, projectID) {
		return nil, fmt.Errorf("usuário não possui o repositório")
	}
	_, err := edb.db.Exec("INSERT INTO PENDING_PROJECT(event_id, project_id) VALUES (?, ?)", eventID, projectID)
	if err != nil {
		return nil, err 
	}
	project, err := pdb.getProjectByID(projectID)
	if err != nil {
		return nil, err 
	}
	fmt.Printf(project.Title)
	return project, nil 
}
func (edb *EventDB) getRoomsByEventID(eventID int64) ([]*entity.RoomEvent, error) {
	rows, err := edb.db.Query(queries.GET_ROOMS_BY_EVENT_ID, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rooms []*entity.RoomEvent
	for rows.Next() {
		var room entity.RoomEvent
		err = rows.Scan(&room.ID, &room.Name, &room.QuantityProjects)
		if err != nil {
			return nil, err
		}
		projects, err := edb.getProjectsByRoomID(room.ID)
		if err != nil {
			return nil, err
		}
		room.Projects = projects
		rooms = append(rooms, &room)
	}
	return rooms, nil
}

func (edb *EventDB) getProjectsByRoomID(roomID int) ([]*entity.Project, error) {
	rows, err := edb.db.Query(queries.GET_REPOSITORIES_BY_ROOM_ID, roomID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var repositories []*entity.Project
	for rows.Next() {
		var repo entity.Project
		err = rows.Scan(&repo.ID, &repo.Title, &repo.Description, &repo.Content, &repo.FirstOwnerID, &repo.FirstOwnerName)
		if err != nil {
			return nil, err
		}
		repositories = append(repositories, &repo)
	}
	return repositories, nil
}

func (edb *EventDB) getEventByID(eventID int64) ([]*entity.Event, error) {
	rows, err := edb.db.Query(queries.GET_EVENT_BY_ID, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []*entity.Event
	for rows.Next() {
		var event entity.Event
		err = rows.Scan(&event.ID, &event.Name, &event.Address, &event.Date, &event.Description, &event.InstitutionID, &event.InstitutionName)
		if err != nil {
			return nil, err
		}
		rooms, err := edb.getRoomsByEventID(event.ID)
		if err != nil {
			return nil, err
		}
		subscribers, err := edb.getSubscribersByEventID(eventID)
		if err != nil {
			return nil, err
		}
		event.Subscribers = subscribers
		event.Rooms = rooms
		events = append(events, &event)
	}
	return events, err
}

func (edb *EventDB) getSubscribersByEventID(eventID int64) ([]*entity.CommonUserBasic, error) {
	rows, err := edb.db.Query(queries.GET_SUBSCRIBERS_BY_EVENT_ID, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subscribers []*entity.CommonUserBasic
	for rows.Next() {
		var subscriber entity.CommonUserBasic
		err = rows.Scan(&subscriber.ID, &subscriber.Name)
		if err != nil {
			return nil, err
		}
		subscribers = append(subscribers, &subscriber)
	}
	return subscribers, nil
}
