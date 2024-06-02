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
	result, err := edb.db.Exec("INSERT INTO EVENT(name, address, date, description, institution_id) FROM VALUES (?, ?, ?, ?, ?)", &event.Name, &event.Address, &event.Date, &event.Description, &institutionID)
	if err != nil {
		return nil, err 
	}
	event.ID, err = result.LastInsertId()
	if err != nil{ 
		return nil, err 
	}
	_, err = edb.db.Exec("INSERT INTO OWNER_EVENT(event_id, owner_id) VALUES (?, ?)", &event.ID, &event.InstitutionID)
	if err != nil {
		return nil, err 
	}
	return event, nil
}