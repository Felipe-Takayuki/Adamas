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
	projectDB := database.NewProjectDB(db)
	institutionDB := database.NewInstitutionDB(db)
	eventDB := database.NewEventDB(db)

	userService := service.NewUserService(*userDB)
	projectService := service.NewProjectService(*projectDB)
	institutionService := service.NewInstitutionService(*institutionDB)
	eventService := service.NewEventService(eventDB)

	webUserService := webserver.NewWebUserHandler(*userService)
	webProjectService := webserver.NewProjectHandler(projectService)
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

	c.Get("/project/user/{user_id}", webProjectService.GetProjectsByUser)
	c.Get("/project/search/{project_title}", webProjectService.GetProjectsByName)
	c.Get("/project/search", webProjectService.GetProjects)

	c.Get("/event/search/{event}", webEventService.GetEventByName)
	c.Get("/event/search", webEventService.GetEvents)
	// Rotas protegidas
	c.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/project", webProjectService.CreateProject)
		r.Put("/project/{project_id}", webProjectService.EditProject)
		r.Delete("/project/{project_id}", webProjectService.DeleteProject)
		r.Post("/event", webEventService.CreateEvent)
		r.Post("/event/{event_id}/room", webEventService.AddRoomInEvent)
		r.Post("/event/{event_id}/subscribe",webEventService.EventRegistration)
		r.Get("/event/{event_id}/subscribers", webEventService.GetSubscribers)
		r.Post("/event/{event_id}/participation",webEventService.EventRequestParticipation)
		r.Post("/event/{event_id}/approve-participation",webEventService.ApproveParticipation)
		r.Post("/project/{project_id}/category", webProjectService.SetCategory)
		r.Post("/project/{project_id}/comment", webProjectService.SetComment)
		r.Delete("/project/{project_id}/comment", webProjectService.DeleteComment)
	},
	)

	return c
}
