package queries

const CREATE_INSTITUTION = "INSERT INTO INSTITUTION_USER(name, email, password, cnpj) VALUES (?, ?, ?, ?)"

const LOGIN_INSTITUTION = "SELECT id, name, email, cnpj FROM INSTITUTION_USER WHERE email = ? and password = ?"