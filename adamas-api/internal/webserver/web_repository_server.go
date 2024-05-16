package webserver

import (
	"encoding/json"
	"net/http"
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/entity"
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/service"
)

type WebRepoHandler struct {
	RepoService *service.RepositoryService
}

func NewRepoHandler(repoService *service.RepositoryService) *WebRepoHandler {
	return &WebRepoHandler{
		RepoService: repoService,
	}
}

func (wph *WebRepoHandler) CreateRepo(w http.ResponseWriter, r *http.Request) {
	var req *entity.RepositoryRequestFirst
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := wph.RepoService.CreateRepo(req.Title, req.Description, req.FirstOwnerID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)

}
