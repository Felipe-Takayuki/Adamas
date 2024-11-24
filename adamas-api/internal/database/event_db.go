package database

import (
	"database/sql"
	"fmt"

	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/entity"
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/utils"
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

func (edb *EventDB) CreateEvent(name, address, startDate, endDate, description string, institutionID int64) (*entity.Event, error) {
	event := entity.NewEvent(name, address, startDate, endDate, description, institutionID)
	result, err := edb.db.Exec(queries.CREATE_EVENT, &event.Name, &event.Address, &event.StartDate, &event.EndDate, &event.Description)
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

func (edb *EventDB) DeleteEvent(eventID int64, email, password string) error {

	institutionID, err := validateInstitution(edb.db, email, password)
	if err != nil {
		return err
	}
	isOwner := isEventOwner(edb.db, eventID, institutionID)
	if !isOwner {
		return fmt.Errorf("a instituição não possui o evento")
	}
	err = deleteRoomsEvent(edb.db, institutionID)
	if err != nil {
		return err
	}
	err = deleteSubscribers(edb.db, eventID)
	if err != nil {
		return err
	}
	err = deleteOwnerEvent(edb.db, institutionID, eventID)
	if err != nil {
		return err
	}
	err = deleteProjectsEvent(edb.db, eventID)
	if err != nil {
		return err
	}
	return nil
}

func (edb *EventDB) EditRoom(roomID, eventID, quantityProjects, ownerID int64, roomName string) (*entity.RoomEvent, error) {
	if !isEventOwner(edb.db, eventID, ownerID) {
		return nil, fmt.Errorf("instituição não possui o evento")
	}

	if roomName != "" {
		_, err := edb.db.Exec(queries.UPDATE_QUANTITY_PROJECTS_ROOM)
		if err != nil {
			return nil, err
		}
	}

	if quantityProjects != 0 {
		_, err := edb.db.Exec(queries.UPDATE_QUANTITY_PROJECTS_ROOM)
		if err != nil {
			return nil, err
		}
	}
	room := &entity.RoomEvent{Name: roomName, QuantityProjects: int(quantityProjects)}
	return room, nil

}

func (edb *EventDB) DeleteRoom(roomID, eventID int64) error {
	_, err := edb.db.Exec(queries.DELETE_ROOM, roomID, eventID)
	if err != nil {
		return err
	}
	return nil
}
func deleteOwnerEvent(db *sql.DB, institutionID, eventID int64) error {
	_, err := db.Exec(queries.DELETE_EVENT_OWNER, institutionID, eventID)
	if err != nil {
		return err
	}
	return nil
}

func deleteProjectsEvent(db *sql.DB, eventID int64) error {
	_, err := db.Exec(queries.DELETE_EVENT_PROJECTS, eventID)
	if err != nil {
		return err
	}
	return nil
}

func deleteRoomsEvent(db *sql.DB, eventID int64) error {
	_, err := db.Exec(queries.DELETE_EVENT_ROOMS, eventID)
	if err != nil {
		return err
	}
	return nil
}
func deleteSubscribers(db *sql.DB, eventID int64) error {
	_, err := db.Exec(queries.DELETE_EVENT_SUBSCRIBERS, eventID)
	if err != nil {
		return err
	}
	return nil
}
func validateInstitution(db *sql.DB, email, password string) (int64, error) {
	var institutionID int64
	err := db.QueryRow(queries.VALIDATE_INSITUTION_USER, email, utils.EncriptKey(password)).Scan(&institutionID)
	if err != nil {
		return 0, err
	}
	return institutionID, nil

}
func (edb *EventDB) AddRoomInEvent(eventID, ownerID int64, roomName string, quantityProjects int) ([]*entity.RoomEvent, error) {
	isOwner := isEventOwner(edb.db, eventID, ownerID)
	if !isOwner {
		return nil, fmt.Errorf("a instituição não possui o evento")
	}
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
		err = rows.Scan(&event.ID, &event.Name, &event.Address, &event.StartDate, &event.EndDate, &event.Description, &event.InstitutionID, &event.InstitutionName)
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
		err = rows.Scan(&event.ID, &event.Name, &event.Address, &event.StartDate, &event.EndDate, &event.Description, &event.InstitutionID, &event.InstitutionName)
		rooms, err := edb.getRoomsByEventID(event.ID)
		if err != nil {
			return nil, err
		}
		subscribers, err := edb.GetSubscribersByEventID(event.ID, event.InstitutionID)
		if err != nil {
			return nil, err
		}
		event.Subscribers = subscribers
		event.Rooms = rooms
		events = append(events, &event)
	}
	return events, err
}

func (edb *EventDB) EventRegistration(eventID, userID int64) (*entity.Event, error) {
	_, err := edb.db.Exec("INSERT INTO SUBSCRIBERS_EVENT(event_id, user_id) VALUES (?, ?)", eventID, userID)
	if err != nil {
		return nil, err
	}
	event, err := edb.GetEventByID(eventID)
	if err != nil {
		return nil, err
	}
	return event, nil
}
func (edb *EventDB) EventRequestParticipation(eventID, userID, projectID int64) (*entity.Project, error) {
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
	return project, nil
}

func (edb *EventDB) GetPendingProjectsInEvent(eventID, ownerID int64) ([]*entity.Project, error) {
	pdb := NewProjectDB(edb.db)
	isOwner := isEventOwner(edb.db, eventID, ownerID)
	if !isOwner {
		return nil, fmt.Errorf("a instituição não possui o evento")
	}
	rows, err := edb.db.Query(queries.GET_PENDING_PROJECTS, eventID)
	if err != nil {
		return nil, err
	}
	
	defer rows.Close()
	var projects []*entity.Project
	for rows.Next() {
		var project entity.Project
		if err := rows.Scan(&project.ID, &project.Title, &project.Description, &project.Content, &project.FirstOwnerID, &project.FirstOwnerName); err != nil {
			return nil, err
		}
		project.Categories, err = pdb.getCategoriesByRepoID(project.ID)
		if err != nil {
			return nil, err
		}
		project.Comments, err = pdb.getCommentsByRepoID(project.ID)
		if err != nil {
			return nil, err
		}
		project.Owners, err = pdb.getOwnersByProjectID(project.ID)
		if err != nil {
			return nil, err
		}
		project.Likes, err = getLikes(pdb.db, project.ID)
		if err != nil {
			return nil, err
		}
		projects = append(projects, &project)
	}

	return projects, nil 
}

func (edb *EventDB) GetProjectsInEvent(eventID int64) ([]*entity.Project, error) {
	pdb := NewProjectDB(edb.db)
	rows, err := edb.db.Query(queries.GET_PROJECTS_EVENT, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var projects []*entity.Project
	for rows.Next() {
		var project entity.Project
		if err := rows.Scan(&project.ID, &project.Title, &project.Description, &project.Content, &project.FirstOwnerID, &project.FirstOwnerName); err != nil {
			return nil, err
		}
		project.Categories, err = pdb.getCategoriesByRepoID(project.ID)
		if err != nil {
			return nil, err
		}
		project.Comments, err = pdb.getCommentsByRepoID(project.ID)
		if err != nil {
			return nil, err
		}
		project.Owners, err = pdb.getOwnersByProjectID(project.ID)
		if err != nil {
			return nil, err
		}
		project.Likes, err = getLikes(pdb.db, project.ID)
		if err != nil {
			return nil, err
		}
		projects = append(projects, &project)
	}

	return projects, nil 
}

func (edb *EventDB) ApproveParticipation(projectID, ownerID, eventID, roomID int64) ([]*entity.Project, error) {
	isOwner := isEventOwner(edb.db, eventID, ownerID)
	if !isOwner {
		return nil, fmt.Errorf("a instituição não possui o evento")
	}
	_, err := edb.db.Exec(queries.APPROVE_PARTICIPATION, eventID, roomID, projectID)
	if err != nil {
		return nil, err
	}
	_, err = edb.db.Exec(queries.DELETE_PENDING_PARTICIPATION, eventID, projectID)
	if err != nil {
		return nil, err
	}

	projects, err := edb.getProjectsByRoomID(roomID)
	if err != nil {
		return nil, err
	}
	return projects, nil
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

func (edb *EventDB) getProjectsByRoomID(roomID int64) ([]*entity.Project, error) {
	rows, err := edb.db.Query(queries.GET_REPOSITORIES_BY_ROOM_ID, roomID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var repositories []*entity.Project
	projectDB := NewProjectDB(edb.db)
	for rows.Next() {
		var repo entity.Project
		err = rows.Scan(&repo.ID, &repo.Title, &repo.Description, &repo.Content, &repo.FirstOwnerID, &repo.FirstOwnerName)
		if err != nil {
			return nil, err
		}
		categories, err := projectDB.getCategoriesByRepoID(repo.ID)
		if err != nil {
			return nil, err
		}
		comments, err := projectDB.getCommentsByRepoID(repo.ID)
		if err != nil {
			return nil, err
		}
		repo.Categories = categories
		repo.Comments = comments
		repositories = append(repositories, &repo)
	}
	return repositories, nil
}

func (edb *EventDB) GetEventByID(eventID int64) (*entity.Event, error) {
	event := &entity.Event{}
	err := edb.db.QueryRow(queries.GET_EVENT_BY_ID, eventID).Scan(&event.ID, &event.Name, &event.Address, &event.StartDate, &event.EndDate, &event.Description, &event.InstitutionID, &event.InstitutionName)
	if err != nil {
		return nil, err
	}
	return event, err
}

func (edb *EventDB) GetSubscribersByEventID(eventID, ownerID int64) ([]*entity.User, error) {

	isOwner := isEventOwner(edb.db, eventID, ownerID)
	if !isOwner {
		return nil, fmt.Errorf("a instituição não possui o evento")
	}

	rows, err := edb.db.Query(queries.GET_SUBSCRIBERS_BY_EVENT_ID, eventID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subscribers []*entity.User
	for rows.Next() {
		var subscriber entity.User
		err = rows.Scan(&subscriber.ID, &subscriber.Name)
		if err != nil {
			return nil, err
		}
		subscribers = append(subscribers, &subscriber)
	}
	return subscribers, nil
}

func (edb *EventDB) EditEvent(eventID, ownerID int64, name, address, startDate, endDate, description string) (*entity.Event, error) {
	if !isEventOwner(edb.db, eventID, ownerID) {
		return nil, fmt.Errorf("instituição não possui o evento")
	}

	if name != "" {
		_, err := edb.db.Exec(queries.UPDATE_NAME_EVENT, name, eventID)
		if err != nil {
			return nil, err
		}
	}

	if address != "" {
		_, err := edb.db.Exec(queries.UPDATE_ADDRESS_EVENT, address, eventID)
		if err != nil {
			return nil, err
		}
	}

	if startDate != "" {
		_, err := edb.db.Exec(queries.UPDATE_START_DATE_EVENT, startDate, eventID)
		if err != nil {
			return nil, err
		}
	}

	if endDate != "" {
		_, err := edb.db.Exec(queries.UPDATE_END_DATE_EVENT, endDate, eventID)
		if err != nil {
			return nil, err
		}
	}

	if description != "" {
		_, err := edb.db.Exec(queries.UPDATE_DESCRIPTION_EVENT, description, eventID)
		if err != nil {
			return nil, err
		}
	}
	event := entity.Event{Name: name, Address: address, StartDate: startDate, EndDate: endDate, Description: description}

	return &event, nil
}
func isEventOwner(db *sql.DB, eventID, ownerID int64) bool {
	var count int
	err := db.QueryRow(queries.CHECK_EVENT_OWNER, ownerID, eventID).Scan(&count)
	if err != nil {
		return false
	}
	return count > 0

}
