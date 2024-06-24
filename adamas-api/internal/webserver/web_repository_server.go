package webserver

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/entity"
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/entity/reqs"
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
	repoName := chi.URLParam(r, "repo")
	w.Header().Set("Content-Type", "application/json")
	if repoName == "" {
		error := utils.ErrorMessage{Message: "title is required"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
	repositories, err := wph.RepoService.GetRepositoriesByName(repoName)
	if err != nil {
		error := utils.ErrorMessage{Message: err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error)
		return

	}
	json.NewEncoder(w).Encode(repositories)

}
func (wph *WebRepoHandler) GetRepositories(w http.ResponseWriter, r *http.Request) {
	repositories, err := wph.RepoService.GetRepositories()
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		error := utils.ErrorMessage{Message: err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error)
		return

	}
	json.NewEncoder(w).Encode(repositories)
}

func (wph *WebRepoHandler) CreateRepo(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	w.Header().Set("Content-Type", "application/json")
	userType, ok := claims["user_type"].(string)
	if !ok {
		http.Error(w, "user_type is not string!", http.StatusInternalServerError)
		return
	}
	if userType == "common_user" {
		flt64, ok := claims["id"].(float64)
		if !ok {
			http.Error(w, "id is not int!", http.StatusInternalServerError)
			return
		}
		userID := flt64
		var req *reqs.RepositoryRequestFirst
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			error := utils.ErrorMessage{Message: err.Error()}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(error)
			return
		}
		result, err := wph.RepoService.CreateRepo(req.Title, req.Description, req.Content, int(userID))
		if err != nil {
			error := utils.ErrorMessage{Message: err.Error()}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(error)
			return
		}
		json.NewEncoder(w).Encode(result)
	} else {
		error := utils.ErrorMessage{Message: "este usuário não possui essa permissão!"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}

}

func (wph *WebRepoHandler) EditRepo(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	w.Header().Set("Content-Type", "application/json")
	userType, ok := claims["user_type"].(string)
	if !ok {
		http.Error(w, "user_type is not string!", http.StatusInternalServerError)
		return
	}
	if userType == "common_user" {
		var req *entity.RepositoryBasic
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			error := utils.ErrorMessage{Message: err.Error()}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(error)
			return
		}
		result, err := wph.RepoService.EditRepo(req.Title, req.Description, req.Content, int64(req.ID))
		if err != nil {
			error := utils.ErrorMessage{Message: err.Error()}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(error)
			return
		}
		json.NewEncoder(w).Encode(result)
	}else {
		error := utils.ErrorMessage{Message: "este usuário não possui essa permissão!"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
}
func (wph *WebRepoHandler) SetCategory(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	w.Header().Set("Content-Type", "application/json")
	userType, ok := claims["user_type"].(string)
	fmt.Println(userType)
	if !ok {
		http.Error(w, "user_type is not string!", http.StatusInternalServerError)
		return
	}
	if userType == "common_user" {
		var reqs *reqs.SetCategoryRequest
		err := json.NewDecoder(r.Body).Decode(&reqs)
		if err != nil {
			error := utils.ErrorMessage{Message: err.Error()}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(error)
			return
		}
		err = wph.RepoService.SetCategory(reqs.CategoryName, reqs.RepositoryID)
		if err != nil {
			error := utils.ErrorMessage{Message: err.Error()}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(error)
			return
		}
		json.NewEncoder(w).Encode(map[string]interface{}{"category": reqs.CategoryName})
	} else {
		error := utils.ErrorMessage{Message: "este usuário não possui essa permissão!"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
}

func (wph *WebRepoHandler) SetComment(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	w.Header().Set("Content-Type", "application/json")
	userType, ok := claims["user_type"].(string)
	fmt.Println(userType)
	if !ok {
		http.Error(w, "user_type is not string!", http.StatusInternalServerError)
		return
	}
	if userType == "common_user" {
		flt64, ok := claims["id"].(float64)
		if !ok {
			http.Error(w, "id is not int!", http.StatusInternalServerError)
			return
		}
		userID := flt64
		var reqs *reqs.SetCommentRequest
		err := json.NewDecoder(r.Body).Decode(&reqs)
		if err != nil {
			error := utils.ErrorMessage{Message: err.Error()}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(error)
			return
		}
		err = wph.RepoService.SetComment(int64(userID), reqs.RepositoryID, reqs.Comment)
		if err != nil {
			error := utils.ErrorMessage{Message: err.Error()}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(error)
			return
		}
		json.NewEncoder(w).Encode(map[string]interface{}{"Comment": reqs.Comment})
	} else {
		error := utils.ErrorMessage{Message: "este usuário não possui essa permissão!"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
}
