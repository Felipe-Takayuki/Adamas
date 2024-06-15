package database

import (
	"database/sql"

	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/entity"
)

type EventDB struct {
	db *sql.DB
}

func NewEventDB(db *sql.DB) *EventDB {
	return &EventDB{
		db: db,
	}
}

func (edb *EventDB) CreateEvent(name, address, date, description string, institutionID int64 ) (*entity.Event, error) {
	event := entity.NewEvent(name, address, date, description, institutionID)
	result, err := edb.db.Exec("INSERT INTO EVENT(name, address, date, description) VALUES (?, ?, ?, ?)", &event.Name, &event.Address, &event.Date, &event.Description)
	if err != nil {
		return nil, err
	}
	event.ID, err = result.LastInsertId()
	if err != nil {
		return nil, err
	}
	_, err = edb.db.Exec("INSERT INTO OWNER_EVENT(event_id, owner_id) VALUES (?, ?)", &event.ID, &event.InstitutionID)
	if err != nil {
		return nil, err
	}
	return event, nil
}

// func (edb *EventDB) GetEventByName(name string) ([]*entity.Event, error) {
// 	query := `SELECT e.id, e.name, e.address, e.date, e.description, o.owner_id, i.name FROM EVENT e
// 	JOIN OWNER_EVENT o ON e.id = o.event_id
// 	JOIN INSTITUTION_USER i ON o.owner_id = i.id
// 	WHERE e.name = ?
// 	`
	
// 	rows, err := edb.db.Query(query, name)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()
// 	var events []*entity.Event
// 	for rows.Next() {
// 		var event entity.Event
// 		err = rows.Scan(&event.ID ,&event.Name, &event.Address, &event.Date, &event.Description, &event.InstitutionID, &event.InstitutionName)
// 		events = append(events, &event)
// 	}
// 	return events, err
// }

func (edb *EventDB) GetEventByName(name string) ([]*entity.Event, error) {
	queryEvent := `
		SELECT e.id, e.name, e.address, e.date, e.description, o.owner_id, i.name 
		FROM EVENT e
		JOIN OWNER_EVENT o ON e.id = o.event_id
		JOIN INSTITUTION_USER i ON o.owner_id = i.id
		WHERE e.name = ?
	`
	rows, err := edb.db.Query(queryEvent, name)
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

func (edb *EventDB) getRoomsByEventID(eventID int64) ([]*entity.RoomEvent, error) {
	queryRooms := `
		SELECT r.id, r.name, r.quantity_repos 
		FROM ROOM_IN_EVENT r 
		WHERE r.event_id = ?
	`
	rows, err := edb.db.Query(queryRooms, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var rooms []*entity.RoomEvent
	for rows.Next() {
		var room entity.RoomEvent
		err = rows.Scan(&room.ID, &room.Name, &room.QuantityRepositories)
		if err != nil {
			return nil, err
		}
		repositories, err := edb.getRepositoriesByRoomID(room.ID)
		if err != nil {
			return nil, err
		}
		room.Repositories = repositories
		rooms = append(rooms, &room)
	}
	return rooms, nil
}

func (edb *EventDB) getRepositoriesByRoomID(roomID int) ([]*entity.Repository, error) {
	queryRepos := `
		SELECT r.id, r.title, r.description, r.content, u.id, u.name 
		FROM REPOSITORY r
		JOIN REPOSITORY_IN_ROOM rr ON r.id = rr.repository_id
		JOIN INSTITUTION_USER u ON r.owner_id = u.id
		WHERE rr.room_id = ?
	`
	rows, err := edb.db.Query(queryRepos, roomID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var repositories []*entity.Repository
	for rows.Next() {
		var repo entity.Repository
		err = rows.Scan(&repo.ID, &repo.Title, &repo.Description, &repo.Content, &repo.FirstOwnerID, &repo.FirstOwnerName)
		if err != nil {
			return nil, err
		}
		repositories = append(repositories, &repo)
	}
	return repositories, nil
}


func (edb *EventDB) GetEvents() ([]*entity.Event, error) {
	query := `SELECT e.id, e.name, e.address, e.date, e.description, o.owner_id, i.name FROM EVENT e
	JOIN OWNER_EVENT o ON e.id = o.event_id
	JOIN INSTITUTION_USER i ON o.owner_id = i.id`
	rows, err := edb.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []*entity.Event
	for rows.Next() {
		var event entity.Event
		err = rows.Scan(&event.ID ,&event.Name, &event.Address, &event.Date, &event.Description, &event.InstitutionID, &event.InstitutionName)
		events = append(events, &event)
	}
	return events, err
}