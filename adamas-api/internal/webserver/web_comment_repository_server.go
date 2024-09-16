package webserver

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/entity"

	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/utils"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
)

func (wph *WebProjectHandler) SetComment(w http.ResponseWriter, r *http.Request) {
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
			http.Error(w, "project_id is not int!", http.StatusInternalServerError)
			return
		}
		var reqs *entity.Comment
		err = json.NewDecoder(r.Body).Decode(&reqs)
		if err != nil {
			error := utils.ErrorMessage{Message: err.Error()}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(error)
			return
		}
		err = wph.ProjectService.SetComment(int64(userID), int64(projectID), reqs.Comment)
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

func (wph *WebProjectHandler) EditComment(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	w.Header().Set("Content-Type", "application/json")
	userID, ok := claims["id"].(float64)
	if !ok {
		http.Error(w, "id is not int!", http.StatusInternalServerError)
		return
	}
	projectID, err := strconv.Atoi(chi.URLParam(r, "project_id"))
	if err != nil {
		http.Error(w, "project_id is not int!", http.StatusInternalServerError)
		return
	}
	var reqs *entity.Comment
	err = json.NewDecoder(r.Body).Decode(&reqs)
	if err != nil {
		error := utils.ErrorMessage{Message: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
	comment,err := wph.ProjectService.EditComment(reqs.Comment,int64(projectID),reqs.CommentID, int64(userID))
	if err != nil {
		error := utils.ErrorMessage{Message: err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error)
		return
	}
	json.NewEncoder(w).Encode(comment) 


}
func (wph *WebProjectHandler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	w.Header().Set("Content-Type", "application/json")
	userType, ok := claims["user_type"].(string)
	if !ok {
		http.Error(w, "user_type is not string!", http.StatusInternalServerError)
		return
	}
	if userType == "common_user" {
		var comment *entity.Comment
		projectID, err := strconv.Atoi(chi.URLParam(r, "project_id"))
		if err != nil {
			http.Error(w, "project_id is not int!", http.StatusInternalServerError)
			return
		}
		err = json.NewDecoder(r.Body).Decode(&comment)
		if err != nil {
			error := utils.ErrorMessage{Message: err.Error()}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(error)
			return
		}
		err = wph.ProjectService.DeleteComment(int64(comment.CommentID), int64(projectID))
		if err != nil {
			error := utils.ErrorMessage{Message: err.Error()}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(error)
			return
		}
		json.NewEncoder(w).Encode(map[string]interface{}{"deleted_comment": projectID})
	} else {
		error := utils.ErrorMessage{Message: "este usuário não possui essa permissão!"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
}
