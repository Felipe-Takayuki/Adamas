package queries 

const GET_REPOSITORY_BY_NAME =`
	SELECT r.id, r.title, r.description, r.content, o.owner_id, u.name FROM REPOSITORY r
	JOIN OWNERS_REPOSITORY o ON r.id = o.repository_id 
	JOIN COMMON_USER u ON o.owner_id = u.id WHERE r.title = ?`

const GET_REPOSITORIES = `
	SELECT r.id, r.title, r.description,r.content, o.owner_id, u.name FROM REPOSITORY r 
	JOIN OWNERS_REPOSITORY o ON r.id = o.repository_id 
	JOIN COMMON_USER u ON o.owner_id = u.id `

const CREATE_REPOSITORY ="INSERT INTO REPOSITORY(title, description,content) VALUES (?,?,?)"

const GET_OWNER_NAME_BY_ID = "SELECT name FROM COMMON_USER WHERE id = ?"

const SET_OWNER = "INSERT INTO OWNERS_REPOSITORY(repository_id, owner_id) VALUES (?, ?)" 

const SET_CATEGORY = "INSERT INTO CATEGORY_REPO(category_id, repository_id) VALUES (?,?)"