package queries

const CREATE_USER = "INSERT INTO COMMON_USER(name, nickname, description, email, password) VALUES( ?, ?, ?, ?, ?)"

const LOGIN_USER = "SELECT id, name, email FROM COMMON_USER WHERE email = ? and password = ?"