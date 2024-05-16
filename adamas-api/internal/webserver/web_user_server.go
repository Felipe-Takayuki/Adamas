package webserver

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/entity"
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/service"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
)

type WebUserHandler struct {
	UserService *service.UserService
}

var tokenString string

func NewWebUserHandler(userService service.UserService) *WebUserHandler {
	return &WebUserHandler{UserService: &userService}
}

func (wub *WebUserHandler) GetRepositoriesByUserName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	if name == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}
	repositories, err := wub.UserService.GetRepositoriesByUserName(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(repositories)

}

func (wub *WebUserHandler) CreateUser(w http.ResponseWriter, r *http.Request, tokenAuth *jwtauth.JWTAuth) {
	var user *entity.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := wub.UserService.CreateUser(user.Name, user.Email, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		claims := map[string]interface{}{"id": result.Id, "name": result.Name, "email": result.Email, "exp" : jwtauth.ExpireIn(time.Minute * 1)}
		_, tokenString, _ = tokenAuth.Encode(claims)
		json.NewEncoder(w).Encode(tokenString)
	}

}

func (wub *WebUserHandler) LoginUser(w http.ResponseWriter, r *http.Request, tokenAuth *jwtauth.JWTAuth) {
	var user entity.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	result, err := wub.UserService.LoginUser(user.Email, user.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	} else {
		claims := map[string]interface{}{"id": result.Id, "name": result.Name, "email": result.Email, "exp" : jwtauth.ExpireIn(time.Minute * 1)}
		_, tokenString, _ = tokenAuth.Encode(claims)
		json.NewEncoder(w).Encode(tokenString)
	}

}
