package queries

const GET_PROJECT_BY_NAME = `
 SELECT DISTINCT p.id, p.title, p.description, p.content, u.id, u.name FROM PROJECT p
 JOIN OWNERS_PROJECT o ON p.id = o.project_id 
 JOIN COMMON_USER u ON u.id = p.first_owner_id WHERE p.title = ?`

const GET_PROJECTS_BY_NAME_CATEGORY = `
 SELECT DISTINCT p.id, p.title, p.description, p.content, u.id, u.name FROM PROJECT p
 JOIN OWNERS_PROJECT o ON p.id = o.project_id 
 JOIN COMMON_USER u ON u.id = p.first_owner_id 
 JOIN CATEGORY_PROJECT cp ON cp.project_id = p.id
 WHERE cp.category_id IN (%s)
 AND p.title = ?
`
const GET_PROJECT_BY_ID = `
 SELECT p.id, p.title, p.description, p.content, o.owner_id, u.name FROM PROJECT p
 JOIN OWNERS_PROJECT o ON p.id = o.project_id 
 JOIN COMMON_USER u ON o.owner_id = u.id WHERE p.id = ?
 `

const GET_PROJECTS_BY_CATEGORIES = `
 SELECT DISTINCT p.id, p.title, p.description, p.content, o.owner_id, u.name FROM PROJECT p
 JOIN OWNERS_PROJECT o ON p.id = o.project_id
 JOIN COMMON_USER u ON u.id = p.first_owner_id
 JOIN CATEGORY_PROJECT cp ON cp.project_id = p.id
 WHERE cp.category_id IN (%s)
`

const GET_PROJECTS = `
 SELECT DISTINCT p.id, p.title, p.description,p.content, u.id, u.name FROM PROJECT p 
 JOIN OWNERS_PROJECT o ON p.id = o.project_id 
 JOIN COMMON_USER u ON u.id = p.first_owner_id`

const GET_PROJECTS_BY_USER = `
 SELECT p.id, p.title, p.description,p.content, o.owner_id, u.name FROM PROJECT p 
 JOIN OWNERS_PROJECT o ON p.id = o.project_id 
 JOIN COMMON_USER u ON o.owner_id = u.id
 WHERE u.id = ?
`
const CREATE_PROJECT = "INSERT INTO PROJECT(title, description,content, first_owner_id) VALUES (?,?,?,?)"

const UPDATE_CONTENT_PROJECT = `
 UPDATE PROJECT 
 SET content = ? 
 WHERE id = ?`
const UPDATE_TITLE_PROJECT = `
 UPDATE PROJECT 
 SET title = ? 
 WHERE id = ?`
const UPDATE_DESCRIPTION_PROJECT = `
 UPDATE PROJECT 
 SET description = ? 
 WHERE id = ?`

const UPDATE_COMMENT = `
 UPDATE COMMENT
 SET comment = ? 
 WHERE id = ?
`
const DELETE_PROJECT = `
 DELETE FROM PROJECT 
 WHERE id = ?
`
const DELETE_OWNER_PROJECT = `
 DELETE FROM OWNERS_PROJECT 
 WHERE owner_id = ? 
 AND project_id = ?
`
const VALIDATE_USER = `
 SELECT id 
 FROM COMMON_USER 
 WHERE email = ? 
 AND password = ?
`

const CHECK_PROJECT_OWNER = `
 SELECT COUNT(*) 
 FROM OWNERS_PROJECT 
 WHERE owner_id = ? 
 AND project_id = ?
`

const CHECK_COMMENT_OWNER = `
 SELECT COUNT(*) 
 FROM COMMENT 
 WHERE owner_id = ? 
 AND id = ?
`

const GET_OWNER_NAME_BY_ID = "SELECT name FROM COMMON_USER WHERE id = ?"

const SET_OWNER = "INSERT INTO OWNERS_PROJECT(project_id, owner_id) VALUES (?, ?)"

const SET_CATEGORY = "INSERT INTO CATEGORY_PROJECT(category_id, project_id) VALUES (?,?)"

const GET_CATEGORIES_BY_PROJECT = `
    SELECT c.name, c.id FROM CATEGORY_PROJECT cp
    JOIN CATEGORY c ON cp.category_id = c.id
	JOIN PROJECT p ON cp.project_id = p.id
	WHERE cp.project_id = ?
`

const GET_OWNERS_BY_PROJECT = `
	SELECT op.owner_id, u.name FROM OWNERS_PROJECT op
	JOIN COMMON_USER u ON op.owner_id = u.id
	WHERE op.project_id = ?
`
const GET_COMMENT_BY_ID = `
 SELECT com.id, u.id, u.name, com.comment FROM COMMENT com
 JOIN PROJECT r ON com.project_id = r.id
 JOIN COMMON_USER u ON com.owner_id = u.id
 WHERE com.id = ?
`
const GET_COMMENTS_BY_PROJECT = `
 SELECT com.id, u.id, u.name, com.comment FROM COMMENT com
 JOIN PROJECT r ON com.project_id = r.id
 JOIN COMMON_USER u ON com.owner_id = u.id
 WHERE com.project_id = ?
`
const SET_COMMENT = "INSERT INTO COMMENT (owner_id, project_id, comment) VALUES (?, ?, ?)"

const DELETE_COMMENT = "DELETE FROM COMMENT WHERE id = ? and project_id = ?"
