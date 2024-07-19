package queries

const CREATE_EVENT = "INSERT INTO EVENT(name, address, date, description) VALUES (?, ?, ?, ?)"

const SET_OWNER_EVENT = "INSERT INTO OWNER_EVENT(event_id, owner_id) VALUES (?, ?)"

const AddRoomInEvent = "INSERT INTO ROOM_IN_EVENT(event_id, name,quantity_projects) VALUES (?, ?, ?)"

const GET_EVENT_BY_NAME = `
	SELECT e.id, e.name, e.address, e.date, e.description, o.owner_id, i.name FROM EVENT e
	JOIN OWNER_EVENT o ON e.id = o.event_id
	JOIN INSTITUTION_USER i ON o.owner_id = i.id
	WHERE e.name = ?`

	const GET_EVENT_BY_ID = `
	SELECT e.id, e.name, e.address, e.date, e.description, o.owner_id, i.name FROM EVENT e
	JOIN OWNER_EVENT o ON e.id = o.event_id
	JOIN INSTITUTION_USER i ON o.owner_id = i.id
	WHERE e.id = ?`

const GET_ROOMS_BY_EVENT_ID = `
	SELECT rie.id, rie.name, rie.quantity_projects FROM ROOM_IN_EVENT rie
	WHERE rie.event_id = ?`

const GET_SUBSCRIBERS_BY_EVENT_ID = `
	SELECT se.user_id, u.name FROM COMMON_USER u 
	JOIN SUBSCRIBERS_EVENT se ON se.user_id = u.id
	WHERE u.id = ?
`

const CHECK_EVENT_OWNER = `
 SELECT COUNT(*) 
 FROM OWNER_EVENT 
 WHERE owner_id = ? 
 AND event_id = ?
`

const GET_REPOSITORIES_BY_ROOM_ID = `
	SELECT p.id, p.title, p.description, p.content, u.id, u.name FROM PROJECT p
	JOIN PROJECT_IN_ROOM pr ON p.id = pr.project_id
	JOIN OWNERS_PROJECT op ON op.project_id = p.id
	JOIN COMMON_USER u ON u.id = op.owner_id
	WHERE pr.room_id = ?`

const GET_EVENTS = `
	SELECT e.id, e.name, e.address, e.date, e.description, o.owner_id, i.name FROM EVENT e
	JOIN OWNER_EVENT o ON e.id = o.event_id
	JOIN INSTITUTION_USER i ON o.owner_id = i.id`