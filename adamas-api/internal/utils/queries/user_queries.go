package queries

const CREATE_USER = "INSERT INTO COMMON_USER(name, nickname, description, email, password) VALUES( ?, ?, ?, ?, ?)"

const LOGIN_USER = "SELECT id, name, email FROM COMMON_USER WHERE email = ? and password = ?"

const GET_USERS = "SELECT id, name FROM COMMON_USER"

const GET_USERS_BY_NAME = "SELECT id, name FROM COMMON_USER WHERE name LIKE ?"