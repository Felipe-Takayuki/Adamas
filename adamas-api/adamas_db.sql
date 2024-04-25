DROP DATABASE IF EXISTS ADAMAS_DB;

CREATE DATABASE ADAMAS_DB;

USE ADAMAS_DB;


CREATE TABLE INSTITUTION_USER(
    id int auto_increment NOT NULL PRIMARY KEY,
    name varchar(255) NOT NULL,
    email varchar(255) NOT NULL UNIQUE,
    password varchar(64) NOT NULL,
    cnpj char(14) NOT NULL UNIQUE
);
CREATE TABLE COMMON_USER(
    id int auto_increment NOT NULL PRIMARY KEY,
    name varchar(255) NOT NULL,
    email varchar(255) NOT NULL UNIQUE,
    institution_id int NULL,
    password varchar(64) NOT NULL,
    FOREIGN KEY (institution_id) REFERENCES INSTITUTION_USER(id)
);

DROP TABLE IF EXISTS EVENT;
CREATE TABLE EVENT(
    id int auto_increment NOT NULL PRIMARY KEY,
    data date NOT NULL,
    name varchar(255) NOT NULL, 
    owner_id int NOT NULL,
    FOREIGN KEY (owner_id) REFERENCES INSTITUTION_USER(id)
); 

CREATE TABLE REPOSITORIES_IN_EVENT(
    repository_id int NOT NULL,
    event_id int NOT NULL,
)


CREATE TABLE REPOSITORY(
    id int auto_increment NOT NULL PRIMARY KEY,
    title varchar(255) NOT NULL,
    description varchar(255) NOT NULL
);
CREATE TABLE OWNERS_REPOSITORY(
    repository_id int NOT NULL REFERENCES REPOSITORY(id),
    owner_id int NOT NULL REFERENCES COMMON_USER(id),
    PRIMARY KEY(repository_id, owner_id)
);
CREATE TABLE BLOC_REPOSITORY(
    id int NOT NULL auto_increment PRIMARY KEY,
    repository_id varchar(36) NOT NULL,
    subtitle varchar(255) NOT NULL,
    content varchar(255) NOT NULL
);

CREATE TABLE CATEGORY_REPO(
    category_id int NOT NULL,
    repository_id int NOT NULL,
    PRIMARY KEY(repository_id, repository_id)
);

CREATE TABLE CATEGORY(
    id int auto_increment NOT NULL PRIMARY KEY,
    name varchar(200) NOT NULL,
); 