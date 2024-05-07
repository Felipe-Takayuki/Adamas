package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/entity"
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/service"
	"github.com/go-chi/jwtauth"
)

type WebUserHandler struct {
	UserService *service.UserService
}
var tokenString string 

func NewWebUserHandler(userService service.UserService) *WebUserHandler {
	return &WebUserHandler{UserService: &userService}
}

func (wub *WebUserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
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
	}
	json.NewEncoder(w).Encode(result)
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
		_, tokenString, _ = tokenAuth.Encode(map[string]interface{}{
			"id": result.Id, "name": result.Name, "email": result.Email})
		json.NewEncoder(w).Encode(tokenString)
	}
	
	
}
