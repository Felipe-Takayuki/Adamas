package webserver

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/entity/reqs"
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/utils"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
)

func (wph *WebRepoHandler) SetComment(w http.ResponseWriter, r *http.Request) {
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

		repositoryID := chi.URLParam(r, "repository_id")
		repoID, err := strconv.Atoi(repositoryID)
		if err != nil {
			http.Error(w, "repository_id is not string!", http.StatusInternalServerError)
			return
		}
		var reqs *reqs.SetCommentRequest
		err = json.NewDecoder(r.Body).Decode(&reqs)
		if err != nil {
			error := utils.ErrorMessage{Message: err.Error()}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(error)
			return
		}
		err = wph.RepoService.SetComment(int64(userID), int64(repoID), reqs.Comment)
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

func (wph *WebRepoHandler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	w.Header().Set("Content-Type", "application/json")
	userType, ok := claims["user_type"].(string)
	if !ok {
		http.Error(w, "user_type is not string!", http.StatusInternalServerError)
		return
	}
	if userType == "common_user" {
		var commentID *reqs.CommentID
		repositoryID := chi.URLParam(r, "repository_id")
		repoID, err := strconv.Atoi(repositoryID)
		if err != nil {
			http.Error(w, "repository_id is not int!", http.StatusInternalServerError)
			return
		}
		err = json.NewDecoder(r.Body).Decode(&commentID)
		if err != nil {
			error := utils.ErrorMessage{Message: err.Error()}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(error)
			return
		}
		err = wph.RepoService.DeleteComment(int64(commentID.CommentID), int64(repoID))
		if err != nil {
			error := utils.ErrorMessage{Message: err.Error()}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(error)
			return
		}
		json.NewEncoder(w).Encode(map[string]interface{}{"deleted_comment": repoID})
	} else {
		error := utils.ErrorMessage{Message: "este usuário não possui essa permissão!"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
}
