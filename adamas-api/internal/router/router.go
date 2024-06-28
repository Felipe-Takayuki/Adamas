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
	repoDB := database.NewRepoDB(db)
	institutionDB := database.NewInstitutionDB(db)
	eventDB := database.NewEventDB(db)

	userService := service.NewUserService(*userDB)
	repoService := service.NewRepoService(*repoDB)
	institutionService := service.NewInstitutionService(*institutionDB)
	eventService := service.NewEventService(eventDB)

	webUserService := webserver.NewWebUserHandler(*userService)
	webRepoService := webserver.NewRepoHandler(repoService)
	webInstitutionService := webserver.NewWebInstiHandler(institutionService)
	webEventService := webserver.NewWebEventHandler(eventService)

	c := chi.NewRouter()
	c.Use(middleware.Logger)
	c.Use(middleware.Recoverer)
	c.Post("/create", func(w http.ResponseWriter, r *http.Request) {
		webUserService.CreateUser(w, r, tokenAuth)
	})
	c.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		webUserService.LoginUser(w, r, tokenAuth)
	})

	c.Post("/create/institution", func(w http.ResponseWriter, r *http.Request) {
		webInstitutionService.CreateInstitution(w, r, tokenAuth)
	})
	c.Post("/login/institution", func(w http.ResponseWriter, r *http.Request) {
		webInstitutionService.LoginInstitution(w, r, tokenAuth)
	})


	c.Get("/repo/{repo}", webRepoService.GetRepositoriesByName)
	c.Get("/repo", webRepoService.GetRepositories)

	c.Get("/event/{event}", webEventService.GetEventByName)
	c.Get("/event", webEventService.GetEvents)
	// Rotas protegidas
	c.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/repo", webRepoService.CreateRepo)
		r.Put("/repo/{repository_id}", webRepoService.EditRepo)
		r.Delete("/repo/{repository_id}", webRepoService.DeleteRepo)
		r.Post("/event", webEventService.CreateEvent)
		r.Post("/repo/{repository_id}/category", webRepoService.SetCategory)
		r.Post("/repo/{repository_id}/comment", webRepoService.SetComment)
		r.Delete("/repo/{repository_id}/comment", webRepoService.DeleteComment)
	},
	)

	return c
}
