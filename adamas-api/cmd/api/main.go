package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/router"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dbHost := os.Getenv("DB_HOST")
	password := "root"
	if dbHost == "" {
		dbHost = "127.0.0.1"
		password = ""
	}
	db, err := sql.Open("mysql", "root:"+password+"@tcp("+dbHost+":3306)/adamas_db")
	if err != nil {
		print("ERROR")
		panic(err.Error())
	}
	c := router.Router(db)
	fmt.Println("Server is running on port 3000")
	http.ListenAndServe(":3000", c)

}
