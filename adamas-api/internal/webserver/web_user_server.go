package webserver

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/entity"
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/service"
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/utils"
	"github.com/go-chi/jwtauth"
)

type WebUserHandler struct {
	UserService *service.UserService
}

var tokenString string

func NewWebUserHandler(userService service.UserService) *WebUserHandler {
	return &WebUserHandler{UserService: &userService}
}

func (wub *WebUserHandler) CreateUser(w http.ResponseWriter, r *http.Request, tokenAuth *jwtauth.JWTAuth) {
	var user *entity.User
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		error := utils.ErrorMessage{Message: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
	result, err := wub.UserService.CreateUser(user.Name, user.Email, user.Password)
	if err != nil {
		error := utils.ErrorMessage{Message: err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error)
		return
	} else {
		claims := map[string]interface{}{"id": result.Id, "name": result.Name, "email": result.Email, "exp": jwtauth.ExpireIn(time.Minute * 10)}
		_, tokenString, _ = tokenAuth.Encode(claims)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"token": tokenString,
		})
	}

}

func (wub *WebUserHandler) LoginUser(w http.ResponseWriter, r *http.Request, tokenAuth *jwtauth.JWTAuth) {
	var user entity.User
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		error := utils.ErrorMessage{Message: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
	result, err := wub.UserService.LoginUser(user.Email, user.Password)
	if err != nil {
		error := utils.ErrorMessage{Message: err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error)
		return
	} else {
		claims := map[string]interface{}{"id": result.Id, "name": result.Name, "email": result.Email, "exp": jwtauth.ExpireIn(time.Minute * 10)}
		_, tokenString, _ = tokenAuth.Encode(claims)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"token": tokenString,
		})
	}

}
