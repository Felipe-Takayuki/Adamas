package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/database"
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/service"
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/webserver"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
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
	userDB := database.NewUserDB(db)
	userService := service.NewUserService(*userDB)

	webUserService := webserver.NewWebUserHandler(*userService)

	c := chi.NewRouter()
	c.Use(middleware.Logger)
	c.Use(middleware.Recoverer)
	c.Post("/create", webUserService.CreateUser)
	c.Post("/login", webUserService.LoginUser)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", c)
}
