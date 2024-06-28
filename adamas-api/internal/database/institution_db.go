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

func (idb *InstitutionDB) CreateInstitution(name, email, password string, cnpj int) (*entity.InstitutionUserExtend, error) {
	institution := entity.NewInstitutionUserExtend(name, email, password, cnpj)
	result, err := idb.db.Exec(queries.CREATE_INSTITUTION, &institution.USER.Name, &institution.USER.Email, &institution.USER.Password, &institution.CNPJ)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err 
	}else {
		institution.USER.ID = id
	}
	return institution, nil
}

func (idb *InstitutionDB) LoginInstitution(email, password string) (*entity.InstitutionUserExtend, error) {
	var institution entity.InstitutionUserExtend
	institution.USER = &entity.User{}
	err := idb.db.QueryRow(queries.LOGIN_INSTITUTION, email, utils.EncriptKey(password)).Scan(
		&institution.USER.ID, &institution.USER.Name, &institution.USER.Email, &institution.CNPJ,
	)
	if err != nil {
		return nil, err
	}
	institution.USER.UserType = "institution_user"
	return &institution, nil
}

