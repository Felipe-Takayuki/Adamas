package database

import (
	"database/sql"

	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/entity"
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/utils"
)

type InstitutionDB struct {
	db *sql.DB
}

func NewInstitutionDB(idb *sql.DB) *InstitutionDB {
	return &InstitutionDB{
		db: idb,
	}
}

func (idb *InstitutionDB) CreateInstitution(name, email, password, cnpj string) (*entity.Institution, error) {
	institution := entity.NewInstitution(name, email, password, cnpj)
	_, err := idb.db.Exec("INSERT INTO INSTITUTION_USER(name, email, password, cnpj) VALUES (?, ?, ?, ?)", &institution.Name, &institution.Email, utils.EncriptKey(institution.Password) ,&institution.CNPJ)
	if err != nil {
		return nil, err
	}
	err = idb.db.QueryRow("SELECT id FROM INSTITUTION_USER WHERE email = ?", email).Scan(&institution.ID)
	if err != nil {
		return nil, err
	}
	return institution, nil 
}

func (idb *InstitutionDB) LoginInstitution(email, password string) (*entity.Institution, error){
	var institution entity.Institution
	err := idb.db.QueryRow("SELECT id, name, email, cnpj FROM INSTITUTION_USER WHERE email = ? and password = ?", email, utils.EncriptKey(password)).Scan(
		institution.ID, institution.Name, institution.Email, institution.CNPJ,
	)
	if err != nil {
		return nil, err
	}
	return &institution, nil
} 