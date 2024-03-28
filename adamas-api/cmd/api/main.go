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
	repositories, err := userService.GetRepositories()
	if err != nil {
		panic(err)
	}
	print(repositories[0].Title)
}
