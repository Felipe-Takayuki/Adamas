package webserver

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/entity"
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/service"
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/utils"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
)

type WebProjectHandler struct {
	ProjectService *service.ProjectService
}

func NewProjectHandler(projectService *service.ProjectService) *WebProjectHandler {
	return &WebProjectHandler{
		ProjectService: projectService,
	}
}

func (wph *WebProjectHandler) GetProjectsByName(w http.ResponseWriter, r *http.Request) {
	projectName := chi.URLParam(r, "project_title")
	w.Header().Set("Content-Type", "application/json")
	if projectName == "" {
		error := utils.ErrorMessage{Message: "title is required"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
	projects, err := wph.ProjectService.GetProjectsByName(projectName)
	if err != nil {
		error := utils.ErrorMessage{Message: err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error)
		return

	}
	json.NewEncoder(w).Encode(projects)

}

func (wph *WebProjectHandler) GetProjectByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	projectID, err := strconv.Atoi(chi.URLParam(r, "project_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	project, err := wph.ProjectService.GetProjectByID(int64(projectID))
	if err != nil {
		error := utils.ErrorMessage{Message: err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error)
		return
	}
	json.NewEncoder(w).Encode(project)

}
func (wph *WebProjectHandler) GetProjects(w http.ResponseWriter, r *http.Request) {
	projects, err := wph.ProjectService.GetProjects()
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		error := utils.ErrorMessage{Message: err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error)
		return

	}
	json.NewEncoder(w).Encode(projects)
}
func (wph *WebProjectHandler) GetProjectsByUser(w http.ResponseWriter, r *http.Request) {
	userID,err := strconv.Atoi(chi.URLParam(r, "user_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	projects, err := wph.ProjectService.GetProjectsByUser(int64(userID))
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		error := utils.ErrorMessage{Message: err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error)
		return

	}
	json.NewEncoder(w).Encode(projects)
}

func (wph *WebProjectHandler) CreateProject(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	w.Header().Set("Content-Type", "application/json")
	userType, ok := claims["user_type"].(string)
	if !ok {
		http.Error(w, "user_type is not string!", http.StatusInternalServerError)
		return
	}
	if userType == "common_user" {
		userID, ok := claims["id"].(float64)
		if !ok {
			http.Error(w, "id is not int!", http.StatusInternalServerError)
			return
		}

		var req *entity.Project
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			error := utils.ErrorMessage{Message: err.Error()}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(error)
			return
		}
		result, err := wph.ProjectService.CreateProject(req.Title, req.Description, req.Content, int(userID))
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
func (wph *WebProjectHandler) DeleteProject(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	w.Header().Set("Content-Type", "application/json")
	userType, ok := claims["user_type"].(string)
	if !ok {
		http.Error(w, "user_type is not string!", http.StatusInternalServerError)
		return
	}

	if userType == "common_user" {
		projectID, err := strconv.Atoi(chi.URLParam(r, "project_id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var reqs *entity.Institution
		err = json.NewDecoder(r.Body).Decode(&reqs)
		if err != nil {
			error := utils.ErrorMessage{Message: err.Error()}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(error)
			return
		}
		err = wph.ProjectService.DeleteProject(reqs.Email, reqs.Password, int64(projectID))
		if err != nil {
			error := utils.ErrorMessage{Message: err.Error()}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(error)
			return
		}
		json.NewEncoder(w).Encode(map[string]interface{}{"deleted_project":projectID})
	}
}

func (wph *WebProjectHandler) EditProject(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	w.Header().Set("Content-Type", "application/json")
	userType, ok := claims["user_type"].(string)
	if !ok {
		http.Error(w, "user_type is not string!", http.StatusInternalServerError)
		return
	}
	if userType == "common_user" {
		userID, ok := claims["id"].(float64)
		if !ok {
			http.Error(w, "id is not int!", http.StatusInternalServerError)
			return
		}
		projectID, err := strconv.Atoi(chi.URLParam(r, "project_id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var req *entity.Project
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			error := utils.ErrorMessage{Message: err.Error()}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(error)
			return
		}
		result, err := wph.ProjectService.EditProject(req.Title, req.Description, req.Content, int64(projectID), int64(userID))
		result.ID = int64(projectID)
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

func (wph *WebProjectHandler) AddNewUserProject(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	w.Header().Set("Content-Type", "application/json")
	userType, ok := claims["user_type"].(string)
	if !ok {
		http.Error(w, "user_type is not string!", http.StatusInternalServerError)
		return
	}
	if userType == "common_user" {
		ownerID, ok := claims["id"].(float64)
		if !ok {
			http.Error(w, "id is not int!", http.StatusInternalServerError)
			return
		}
		projectID, err := strconv.Atoi(chi.URLParam(r, "project_id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var req *entity.User
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			error := utils.ErrorMessage{Message: err.Error()}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(error)
			return
		}
		result, err := wph.ProjectService.AddNewUserProject(int64(projectID), req.ID, int64(ownerID))
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
func (wph *WebProjectHandler) SetCategory(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	w.Header().Set("Content-Type", "application/json")
	userType, ok := claims["user_type"].(string)
	if !ok {
		http.Error(w, "user_type is not string!", http.StatusInternalServerError)
		return
	}
	if userType == "common_user" {
		var reqs *entity.Category
		err := json.NewDecoder(r.Body).Decode(&reqs)
		if err != nil {
			error := utils.ErrorMessage{Message: err.Error()}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(error)
			return
		}
		projectID, err := strconv.Atoi(chi.URLParam(r, "project_id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = wph.ProjectService.SetCategory(reqs.Name, int64(projectID))
		if err != nil {
			error := utils.ErrorMessage{Message: err.Error()}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(error)
			return
		}
		json.NewEncoder(w).Encode(map[string]interface{}{"category": reqs.Name})
	} else {
		error := utils.ErrorMessage{Message: "este usuário não possui essa permissão!"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
}


func (wph *WebProjectHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	w.Header().Set("Content-Type", "application/json")
	userType, ok := claims["user_type"].(string)
	if !ok {
		http.Error(w, "user_type is not string!", http.StatusInternalServerError)
		return
	}
	if userType == "common_user"{
		var reqs *entity.Category
		err := json.NewDecoder(r.Body).Decode(&reqs)
		if err != nil {
			error := utils.ErrorMessage{Message: err.Error()}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(error)
			return
		}
		projectID, err := strconv.Atoi(chi.URLParam(r, "project_id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		ownerID, ok := claims["id"].(float64)
		if !ok {
			http.Error(w, "id is not int!", http.StatusInternalServerError)
			return
		}
		err = wph.ProjectService.DeleteCategory(int64(projectID),int64(ownerID), reqs.Name)
		if err != nil {
			error := utils.ErrorMessage{Message: err.Error()}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(error)
			return
		}
		json.NewEncoder(w).Encode(map[string]interface{}{"category": reqs.Name})
	} else {
		error := utils.ErrorMessage{Message: "este usuário não possui essa permissão!"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
}
