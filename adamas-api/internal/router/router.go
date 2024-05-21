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

	repoDB := database.NewRepoDB(db)
	repoService := service.NewRepoService(*repoDB)
	webRepoService := webserver.NewRepoHandler(repoService)

	institutionDB := database.NewInstitutionDB(db)
	institutionService := service.NewInstitutionService(*institutionDB)
	webInstitutionService := webserver.NewWebInstiHandler(institutionService)

	c := chi.NewRouter()
	c.Use(middleware.Logger)
	c.Use(middleware.Recoverer)
	c.Post("/create", func(w http.ResponseWriter, r *http.Request) {
		webUserService.CreateUser(w, r, tokenAuth)
	})
	c.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		webUserService.LoginUser(w, r, tokenAuth)
	})

	c.Post("create_institution", func(w http.ResponseWriter, r *http.Request) {
		webInstitutionService.CreateInstitution(w, r, tokenAuth)
	})
	c.Post("login_institution", func(w http.ResponseWriter, r *http.Request) {
		
	})


	c.Get("/search/{repo}", webRepoService.GetRepositoriesByName)

	// Rotas protegidas
	c.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/repo", webRepoService.CreateRepo)
	},
	)

	return c
}
