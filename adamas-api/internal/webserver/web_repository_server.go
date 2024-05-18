package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/entity"
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/service"
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/utils"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
)

type WebRepoHandler struct {
	RepoService *service.RepositoryService
}

func NewRepoHandler(repoService *service.RepositoryService) *WebRepoHandler {
	return &WebRepoHandler{
		RepoService: repoService,
	}
}

func (wph *WebRepoHandler) GetRepositoriesByName(w http.ResponseWriter, r *http.Request) {
	repository := chi.URLParam(r, "repo")
	if repository == "" {
		error := utils.ErrorMessage{Message: "id is required"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
	repositories, err := wph.RepoService.GetRepositoriesByName(repository)
	if err != nil {
		error := utils.ErrorMessage{Message: err.Error()}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error)
		return

	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":          repositories[0].ID,
		"title":       repositories[0].Title,
		"description": repositories[0].Description,
	})

}

func (wph *WebRepoHandler) CreateRepo(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())

	flt64, ok := claims["id"].(float64)
	if !ok {
		http.Error(w, "id is not int!", http.StatusInternalServerError)
		return
	}
	userID := flt64
	var req *entity.RepositoryRequestFirst
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		error := utils.ErrorMessage{Message: err.Error()}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
	result, err := wph.RepoService.CreateRepo(req.Title, req.Description, int(userID))
	if err != nil {
		error := utils.ErrorMessage{Message: err.Error()}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error)
		return
	}
	json.NewEncoder(w).Encode(result)

}
