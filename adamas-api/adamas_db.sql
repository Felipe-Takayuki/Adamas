
CREATE DATABASE IF NOT EXISTS adamas_db;

USE adamas_db;


CREATE TABLE IF NOT EXISTS INSTITUTION_USER(
  id int auto_increment NOT NULL PRIMARY KEY,
  name varchar(255) NOT NULL UNIQUE,
  email varchar(255) NOT NULL UNIQUE,
  password varchar(64) NOT NULL,
  cnpj char(14) NOT NULL
);

CREATE TABLE IF NOT EXISTS COMMON_USER(
  id int auto_increment NOT NULL PRIMARY KEY,
  name varchar(255) NOT NULL UNIQUE,
  nickname varchar(20) NOT NULL,
  description varchar(255),
  email varchar(255) NOT NULL UNIQUE,
  institution_id int NULL,
  password varchar(64) NOT NULL,
  FOREIGN KEY (institution_id) REFERENCES INSTITUTION_USER(id)
);


CREATE TABLE IF NOT EXISTS PROJECT(
  id int auto_increment NOT NULL PRIMARY KEY,
  title varchar(255) NOT NULL,
  first_owner_id int NOT NULL,
  description varchar(255) NOT NULL,
  content MEDIUMTEXT NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (first_owner_id) REFERENCES COMMON_USER(id)
);

CREATE TABLE IF NOT EXISTS LIKE_PROJECT(
  project_id int NOT NULL REFERENCES PROJECT(id),
  user_id int NOT NULL REFERENCES COMMON_USER(id),
  PRIMARY KEY(project_id, user_id)
);

CREATE TABLE IF NOT EXISTS OWNERS_PROJECT(
  project_id int NOT NULL REFERENCES PROJECT(id),
  owner_id int NOT NULL REFERENCES COMMON_USER(id),
  PRIMARY KEY(project_id, owner_id)
);

CREATE TABLE IF NOT EXISTS EVENT(
  id int auto_increment NOT NULL PRIMARY KEY,
  name varchar(100) NOT NULL,
  address varchar(255) NOT NULL,
  start_date TIMESTAMP NOT NULL,
  end_date  TIMESTAMP NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  description varchar(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS OWNER_EVENT(
  event_id int NOT NULL REFERENCES EVENT(id),
  owner_id int NOT NULL REFERENCES INSTITUTION_USER(id),
  PRIMARY KEY(event_id, owner_id)
);

CREATE TABLE IF NOT EXISTS SUBSCRIBERS_EVENT(
  event_id int NOT NULL REFERENCES EVENT(id),
  user_id  int NOT NULL REFERENCES COMMON_USER(id),
  PRIMARY KEY(event_id, user_id)
);

CREATE TABLE IF NOT EXISTS ROOM_IN_EVENT(
  id int auto_increment NOT NULL PRIMARY KEY,
  event_id int NOT NULL,
  name varchar(50) NOT NULL UNIQUE,
  quantity_projects int NOT NULL,
  FOREIGN KEY (event_id) REFERENCES EVENT(id) 
);

CREATE TABLE IF NOT EXISTS PENDING_PROJECT(
  event_id int NOT NULL REFERENCES EVENT(id),
  project_id int NOT NULL REFERENCES PROJECT(id),
  PRIMARY KEY(event_id, project_id)
);

CREATE TABLE IF NOT EXISTS PROJECT_IN_ROOM(
  event_id int NOT NULL REFERENCES EVENT(id),
  room_id int NOT NULL REFERENCES ROOM_IN_EVENT(id),
  project_id int NOT NULL REFERENCES PROJECT(id),
  PRIMARY KEY(event_id, room_id, project_id)
);


DELIMITER $$

CREATE TRIGGER IF NOT EXISTS before_insert_project_in_room
BEFORE INSERT ON PROJECT_IN_ROOM
FOR EACH ROW
BEGIN
  DECLARE max_projects INT;
  DECLARE current_projects INT;

  -- Obter o valor de quantity_projects da tabela ROOM_IN_EVENT
  SELECT quantity_projects INTO max_projects
  FROM ROOM_IN_EVENT
  WHERE id = NEW.room_id;

  -- Contar o número atual de projetos na sala
  SELECT COUNT(*) INTO current_projects
  FROM PROJECT_IN_ROOM
  WHERE room_id = NEW.room_id;

  -- Verificar se o limite foi atingido
  IF current_projects >= max_projects THEN
    SIGNAL SQLSTATE '45000'
    SET MESSAGE_TEXT = 'O número máximo de projetos nesta sala foi atingido';
  END IF;
END$$

DELIMITER ;




CREATE TABLE IF NOT EXISTS CATEGORY(
  id int auto_increment NOT NULL PRIMARY KEY,
  name varchar(200) NOT NULL
); 

INSERT INTO CATEGORY(name) values 
("Saúde"),
("Agricultura"),
("Ferramenta"),
("Música"),
("TI"),
("Marketing"),
("Mecânica");

CREATE TABLE IF NOT EXISTS CATEGORY_PROJECT(
  category_id int NOT NULL REFERENCES CATEGORY(id),
  project_id int NOT NULL REFERENCES PROJECT(id),
  PRIMARY KEY(category_id,project_id)
);

DELIMITER $$
CREATE TRIGGER limit_category_count
BEFORE INSERT ON CATEGORY_PROJECT
FOR EACH ROW
BEGIN
  DECLARE category_count INT;
  SELECT COUNT(*) INTO category_count 
  FROM CATEGORY_PROJECT 
  WHERE project_id = NEW.project_id;
  IF category_count >= 3 THEN
    SIGNAL SQLSTATE '45000' 
    SET MESSAGE_TEXT = 'Limite de 3 categorias por projeto atingido';
  END IF;
END$$
DELIMITER ;



CREATE TABLE IF NOT EXISTS COMMENT(
  id int auto_increment NOT NULL PRIMARY KEY,
  owner_id int NOT NULL,
  project_id int NOT NULL,
  comment varchar(255) NOT NULL,
  commented_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (owner_id) REFERENCES COMMON_USER(id),
  FOREIGN KEY (project_id) REFERENCES PROJECT(id)
);
