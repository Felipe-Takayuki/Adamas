package queries

const CREATE_EVENT = "INSERT INTO EVENT(name, address, date, description) VALUES (?, ?, ?, ?)"

const SET_OWNER_EVENT = "INSERT INTO OWNER_EVENT(event_id, owner_id) VALUES (?, ?)"

const GET_EVENT_BY_NAME = `
	SELECT e.id, e.name, e.address, e.date, e.description, o.owner_id, i.name FROM EVENT e
	JOIN OWNER_EVENT o ON e.id = o.event_id
	JOIN INSTITUTION_USER i ON o.owner_id = i.id
	WHERE e.name = ?`

const GET_ROOMS_BY_EVENT_ID = `
	SELECT r.id, r.name, r.quantity_repos FROM ROOM_IN_EVENT r 
	WHERE r.event_id = ?`

const GET_REPOSITORIES_BY_ROOM_ID = `
	SELECT r.id, r.title, r.description, r.content, u.id, u.name FROM REPOSITORY r
	JOIN REPOSITORY_IN_ROOM rr ON r.id = rr.repository_id
	JOIN INSTITUTION_USER u ON r.owner_id = u.id
	WHERE rr.room_id = ?`

const GET_EVENTS = `
	SELECT e.id, e.name, e.address, e.date, e.description, o.owner_id, i.name FROM EVENT e
	JOIN OWNER_EVENT o ON e.id = o.event_id
	JOIN INSTITUTION_USER i ON o.owner_id = i.id`