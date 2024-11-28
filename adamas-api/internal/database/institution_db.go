package database

import (
	"database/sql"

	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/entity"
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/utils"
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/utils/queries"
)

type InstitutionDB struct {
	db *sql.DB
}

func NewInstitutionDB(idb *sql.DB) *InstitutionDB {
	return &InstitutionDB{
		db: idb,
	}
}

func (idb *InstitutionDB) CreateInstitution(name, email, password string, cnpj string) (*entity.Institution, error) {
	institution := entity.NewInstitution(name, email, password, cnpj)
	result, err := idb.db.Exec(queries.CREATE_INSTITUTION, &institution.Name, &institution.Email, &institution.Password, &institution.CNPJ)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err 
	}else {
		institution.ID = id
	}
	return institution, nil
}

func (idb *InstitutionDB) LoginInstitution(email, password string) (*entity.Institution, error) {
	var institution entity.Institution
	err := idb.db.QueryRow(queries.LOGIN_INSTITUTION, email, utils.EncriptKey(password)).Scan(
		&institution.ID, &institution.Name, &institution.Email, &institution.CNPJ,
	)
	if err != nil {
		return nil, err
	}
	institution.UserType = "institution_user"
	return &institution, nil
}

func (idb *InstitutionDB) GetInstitutionByID(institutionID int64)(*entity.Institution, error) {
	institution := &entity.Institution{}

	err := idb.db.QueryRow(queries.GET_INSTITUTION_BY_ID, institutionID).Scan(&institution.ID, &institution.Name)
	if err != nil{
		return nil, err 
	}
	return institution, nil 
}