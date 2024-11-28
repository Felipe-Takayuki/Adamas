package queries

const CREATE_USER = "INSERT INTO COMMON_USER(name, nickname, email, password) VALUES( ?, ?, ?, ?)"

const LOGIN_USER = "SELECT id, name, email FROM COMMON_USER WHERE email = ? and password = ?"

const GET_USERS = "SELECT id, name, nickname, COALESCE(description, '') AS description FROM COMMON_USER"

const GET_USERS_BY_NAME = "SELECT id, name, nickname,  COALESCE(description, '') FROM COMMON_USER WHERE name LIKE ?"

const GET_USER = "SELECT id, name, nickname, COALESCE(description, '') AS description FROM COMMON_USER WHERE id = ?"

const UPDATE_USERNAME = "UPDATE COMMON_USER SET name = ? WHERE id = ?"

const UPDATE_NICKNAME = "UPDATE COMMON_USER SET nickname = ? WHERE id = ?"

const UPDATE_DESCRIPTION = "UPDATE COMMON_USER SET description = ? WHERE id = ?"