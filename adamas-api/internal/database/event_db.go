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

func (edb *EventDB) CreateEvent(name, address, date, description string, institutionID int) (*entity.Event, error) {
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

func (edb *EventDB) GetEventByName(name string) ([]*entity.Event, error) {
	query := `SELECT e.id, e.name, e.address, e.date, e.description, o.owner_id, i.name FROM EVENT e
	JOIN OWNER_EVENT o ON e.id = o.event_id
	JOIN INSTITUTION_USER i ON o.owner_id = i.id
	WHERE e.name = ?
	`
	
	rows, err := edb.db.Query(query, name)
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
