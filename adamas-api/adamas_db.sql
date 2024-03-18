CREATE DATABASE ADAMAS_DB;

USE ADAMAS_DB;

CREATE TABLE COMMON_USER(
    id varchar(36) NOT NULL PRIMARY KEY,
    name varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    password varchar(64) NOT NULL
);
CREATE TABLE INSTUTION_USER(
    id varchar(36) NOT NULL PRIMARY KEY,
    name varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    password varchar(64) NOT NULL,
    cnpj char(14) NOT NULL
);


CREATE TABLE EVENT(
    id varchar(36) NOT NULL PRIMARY KEY,
    data date NOT NULL,
    name varchar(255) NOT NULL, 
    owner_id vachar(36) NOT NULL,
    FOREIGN KEY (owner_id) REFERENCES INSTITUTION_USER(id)
);

CREATE TABLE REPOSITORIES_IN_EVENT(
    
    repository_id varchar(36) NOT NULL,
    event_id varchar(36) NOT NULL,


)




CREATE TABLE REPOSITORY(
    id varchar(36) NOT NULL PRIMARY KEY,
    title varchar(255) NOT NULL,
    description varchar(255) NOT NULL
);
CREATE TABLE OWNERS_REPOSITORY(
    repository_id varchar(36) NOT NULL REFERENCES REPOSITORY(id),
    owner_id varchar(36) NOT NULL REFERENCES COMMON_USER(id),
    PRIMARY KEY(repository_id, owner_id)
);
CREATE TABLE BLOC_REPOSITORY(
    id int NOT NULL auto_increment PRIMARY KEY,
    repository_id varchar(36) NOT NULL,
    subtitle varchar(255) NOT NULL,
    content varchar(255) NOT NULL
);
CREATE TABLE CATEGORY(
    id varchar(36) NOT NULL PRIMARY KEY,
    name varchar(200) NOT NULL,
    repository_id varchar(36) NOT NULL,
    FOREIGN KEY (repository_id) REFERENCES REPOSITORY(id)
); 