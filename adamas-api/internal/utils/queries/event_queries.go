package queries

const CREATE_EVENT = "INSERT INTO EVENT(name, address, start_date, end_date, description) VALUES (?, ?, ?, ?, ?)"

const SET_OWNER_EVENT = "INSERT INTO OWNER_EVENT(event_id, owner_id) VALUES (?, ?)"

const AddRoomInEvent = "INSERT INTO ROOM_IN_EVENT(event_id, name,quantity_projects) VALUES (?, ?, ?)"

const GET_EVENT_BY_NAME = `
	SELECT e.id, e.name, e.address, e.start_date, e.end_date, e.description, o.owner_id, i.name FROM EVENT e
	JOIN OWNER_EVENT o ON e.id = o.event_id
	JOIN INSTITUTION_USER i ON o.owner_id = i.id
	WHERE e.name = ?`

const GET_EVENT_BY_ID = `
	SELECT e.id, e.name, e.address, e.start_date, e.end_date , e.description, o.owner_id, i.name FROM EVENT e
	JOIN OWNER_EVENT o ON e.id = o.event_id
	JOIN INSTITUTION_USER i ON o.owner_id = i.id
	WHERE e.id = ?`

const GET_ROOMS_BY_EVENT_ID = `
	SELECT rie.id, rie.name, rie.quantity_projects FROM ROOM_IN_EVENT rie
	WHERE rie.event_id = ?`

const GET_SUBSCRIBERS_BY_EVENT_ID = `
	SELECT se.user_id, u.name FROM COMMON_USER u 
	JOIN SUBSCRIBERS_EVENT se ON se.user_id = u.id
	WHERE se.event_id = ?
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

const GET_PENDING_PROJECTS = `
	SELECT p.id, p.title, p.description, p.content, u.id, u.name FROM PROJECT p
	JOIN PENDING_PROJECT pp ON p.id = pp.project_id
	JOIN OWNERS_PROJECT op ON op.project_id = p.id
	JOIN COMMON_USER u ON u.id = op.owner_id
	WHERE pp.event_id = ? 
`

const GET_PROJECTS_EVENT = `
	SELECT p.id, p.title, p.description, p.content, u.id, u.name FROM PROJECT p
	JOIN PROJECT_IN_ROOM pir ON p.id = pir.project_id
	JOIN OWNERS_PROJECT op ON op.project_id = p.id
	JOIN COMMON_USER u ON u.id = op.owner_id
	where pir.event_id = ?
`

const GET_EVENTS = `
	SELECT e.id, e.name, e.address, e.start_date, e.end_date, e.description, o.owner_id, i.name FROM EVENT e
	JOIN OWNER_EVENT o ON e.id = o.event_id
	JOIN INSTITUTION_USER i ON o.owner_id = i.id`

const APPROVE_PARTICIPATION = "INSERT INTO PROJECT_IN_ROOM(event_id, room_id, project_id) VALUES (?, ?, ?)"
const DELETE_PENDING_PARTICIPATION = "DELETE FROM PENDING_PROJECT WHERE event_id = ? AND project_id = ?"

const UPDATE_NAME_EVENT = `
 UPDATE EVENT 
 SET name = ? 
 WHERE id = ?`

 const UPDATE_ADDRESS_EVENT = `
 UPDATE EVENT 
 SET address = ? 
 WHERE id = ?
 `

const UPDATE_START_DATE_EVENT = `
UPDATE EVENT 
SET start_date = ? 
WHERE id = ?
`

const UPDATE_END_DATE_EVENT = `
UPDATE EVENT 
SET end_date = ? 
WHERE id = ?
`

const UPDATE_DESCRIPTION_EVENT = `
UPDATE EVENT 
SET description = ? 
WHERE id = ?
`

const UPDATE_ROOM_NAME  = `
UPDATE ROOM_IN_EVENT
SET name = ?
WHERE id = ?
AND event_id =?
`
const UPDATE_QUANTITY_PROJECTS_ROOM  = `
UPDATE ROOM_IN_EVENT
SET name = ?
WHERE id = ?
AND event_id =?
`

const DELETE_PROJECTS_ROOM = `
DELETE FROM PROJECT_IN_ROOM
WHERE room_id = ?`

const DELETE_EVENT_ROOMS = `
DELETE FROM ROOM_IN_EVENT
WHERE event_id = ?
`

const DELETE_ROOM = `
DELETE FROM ROOM_IN_EVENT 
WHERE id = ? and event_id = ? 
`
const DELETE_EVENT_SUBSCRIBERS = `
DELETE FROM SUBSCRIBERS_EVENT 
WHERE event_id = ?`

const DELETE_EVENT_PROJECTS = `
DELETE FROM PROJECT_IN_ROOM
WHERE event_id = ?`