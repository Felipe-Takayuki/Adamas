package queries

const CREATE_USER = "INSERT INTO COMMON_USER(name, nickname, email, password) VALUES( ?, ?, ?, ?)"

const LOGIN_USER = "SELECT id, name, email FROM COMMON_USER WHERE email = ? and password = ?"

const GET_USERS = "SELECT id, name FROM COMMON_USER"

const GET_USERS_BY_NAME = "SELECT id, name FROM COMMON_USER WHERE name LIKE ?"

const UPDATE_USERNAME = "UPDATE COMMON_USER SET name = ? WHERE id = ?"

const UPDATE_NICKNAME = "UPDATE COMMON_USER SET nickname = ? WHERE id = ?"

const UPDATE_DESCRIPTION = "UPDATE COMMON_USER SET description = ? WHERE id = ?"