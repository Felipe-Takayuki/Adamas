package service

import (
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/database"
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/entity"
)

type InstitutionService struct {
	InstitutionDB database.InstitutionDB
}

func NewInstitutionService(institutionDB database.InstitutionDB) *InstitutionService {
	return &InstitutionService{
		InstitutionDB: institutionDB,
	}
}

func (is *InstitutionService) CreateInstitution (name, email, password string, cnpj string) (*entity.Institution, error) {
	institution, err := is.InstitutionDB.CreateInstitution(name, email, password, cnpj)
	if err != nil {
		return nil, err 
	}
	return institution, nil
}

func (is *InstitutionService) LoginInstitution (email, password string) (*entity.Institution, error) {
	institution, err := is.InstitutionDB.LoginInstitution(email, password) 
	if err != nil {
		return nil, err 
	}
	return institution, nil
}

func (is *InstitutionService) GetInstitutionByID (institutionID int64) (*entity.Institution, error) {
	institution,err := is.InstitutionDB.GetInstitutionByID(institutionID)
	if err != nil {
		return nil, err 
	}
	return institution, nil 
}