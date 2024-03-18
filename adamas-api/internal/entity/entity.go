package entity

type User struct {
	Id           string
	Name         string
	Email        string
	Password     string
	Repositories []*Repository
}

type Institution struct {
	Name     string
	Email    string
	Password string
	CNPJ     string
	Events   []string
}

