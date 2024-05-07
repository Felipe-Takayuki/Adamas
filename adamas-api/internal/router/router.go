package router

import (
	"database/sql"
	"net/http"

	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/database"
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/service"
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/webserver"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
)

func Router(db *sql.DB) http.Handler {
	tokenAuth := jwtauth.New("HS256", []byte("secret"), nil)
	userDB := database.NewUserDB(db)
	userService := service.NewUserService(*userDB)
	webUserService := webserver.NewWebUserHandler(*userService)
	c := chi.NewRouter()
	c.Use(middleware.Logger)
	c.Use(middleware.Recoverer)
	c.Post("/create", webUserService.CreateUser)
	c.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		webUserService.LoginUser(w, r, tokenAuth)
	})
	return c
}
