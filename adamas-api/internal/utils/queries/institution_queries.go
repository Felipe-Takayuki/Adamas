package queries

const CREATE_INSTITUTION = "INSERT INTO INSTITUTION_USER(name, email, password, cnpj) VALUES (?, ?, ?, ?)"

const LOGIN_INSTITUTION = "SELECT id, name, email, cnpj FROM INSTITUTION_USER WHERE email = ? and password = ?"

const GET_INSTITUTION_BY_ID = "SELECT id, name FROM INSTITUTION_USER WHERE id = ? "

const VALIDATE_INSTITUTION = `
 SELECT id 
 FROM INSTITUTION_USER 
 WHERE email = ? 
 AND password = ?
`

const DELETE_EVENT_OWNER = ` 
 DELETE FROM OWNER_EVENT 
 WHERE owner_id = ? 
 AND event_id = ?`

 const VALIDATE_INSITUTION_USER = `
 SELECT id 
 FROM INSTITUTION_USER 
 WHERE email = ? 
 AND password = ?
`
