package router

import (
	"database/sql"
	"fmt"
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
	c.Post("/create", func(w http.ResponseWriter, r *http.Request) {
		webUserService.CreateUser(w, r, tokenAuth)
	})
	c.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		webUserService.LoginUser(w, r, tokenAuth)
	})
	c.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Get("/admin", func(w http.ResponseWriter, r *http.Request) {
			_, claims, _ := jwtauth.FromContext(r.Context())
			w.Write([]byte(fmt.Sprintf("protected area. hi %v", claims["name"])))
		})
	},
	)
	return c
}
