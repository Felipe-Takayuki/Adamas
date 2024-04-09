package main

import (
	"database/sql"

	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/database"
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/service"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/adamas_db")
	if err != nil {
		print("ERROR")
		panic(err.Error())
	}
	userDB := database.NewUserDB(db)
	userService := service.NewUserService(*userDB)
	usr, err := userService.CreateUser("Felipe", "felipe@gmail.com", "12345678")
	if err != nil {
		panic(err)
	}
	print(usr.Name, " ", usr.Email, " ", usr.Password)
}
