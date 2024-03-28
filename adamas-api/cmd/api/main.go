package main

import (
	"database/sql"

	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/database"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/adamas_db")
	if err != nil {
		print("ERROR")
		panic(err.Error())
	}
	userDB := database.NewUserDB(db)
	repositories, err := userDB.GetRepositories()
	if err != nil {
		panic(err)
	}
	print(repositories[0].ID)
}
