package webserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/entity"
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/service"
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/utils"
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
	result, err := wub.UserService.CreateUser(user.Name, user.NickName, user.Email, user.Password)
	if err != nil {
		error := utils.ErrorMessage{Message: err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error)
		return
	} else {

		claims := map[string]interface{}{"id": result.ID, "name": result.Name, "email": result.Email, "user_type": result.UserType, "exp": jwtauth.ExpireIn(time.Hour)}
		fmt.Println(result.UserType)
		_, tokenString, _ = tokenAuth.Encode(claims)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"token": tokenString,
		})
	}

}

func (wub *WebUserHandler) LoginUser(w http.ResponseWriter, r *http.Request, tokenAuth *jwtauth.JWTAuth) {
	var login *entity.User
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		error := utils.ErrorMessage{Message: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
	result, err := wub.UserService.LoginUser(login.Email, login.Password)
	if err != nil {
		error := utils.ErrorMessage{Message: err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error)
		return
	} else {
		claims := map[string]interface{}{"id": result.ID, "name": result.Name, "email": result.Email, "user_type": result.UserType, "exp": jwtauth.ExpireIn(time.Hour)}
		_, tokenString, _ = tokenAuth.Encode(claims)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"token": tokenString,
		})
	}

}

func (wub *WebUserHandler) EditUser(w http.ResponseWriter, r *http.Request) {
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
		var edit *entity.User
		err := json.NewDecoder(r.Body).Decode(&edit)
		if err != nil {
			error := utils.ErrorMessage{Message: err.Error()}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(error)
			return
		}
		if userID == float64(edit.ID) {
			user, err := wub.UserService.EditUser(edit.Name, edit.NickName, edit.Description, edit.ID)
			if err != nil {
				error := utils.ErrorMessage{Message: err.Error()}
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(error)
				return
			}
			json.NewEncoder(w).Encode(user)
		} else {
			error := utils.ErrorMessage{Message: "este usuário não possui essa permissão!"}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(error)
			return
		}
	} else {
		error := utils.ErrorMessage{Message: "este usuário não possui essa permissão!"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
}
func (wub *WebUserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users, err := wub.UserService.GetUsers()
	if err != nil {
		error := utils.ErrorMessage{Message: err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error)
		return
	}
	json.NewEncoder(w).Encode(users)
}

func (wub *WebUserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
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
		user, err := wub.UserService.GetUserByID(int64(userID))
		if err != nil {
			error := utils.ErrorMessage{Message: err.Error()}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(error)
			return
		}
		json.NewEncoder(w).Encode(user)
	} else {
		error := utils.ErrorMessage{Message: "este usuário não possui essa permissão!"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}

}
func (wub *WebUserHandler) GetUsersByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	userName := chi.URLParam(r, "username")
	users, err := wub.UserService.GetUsersByName(userName)
	if err != nil {
		error := utils.ErrorMessage{Message: err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error)
		return
	}
	json.NewEncoder(w).Encode(users)

}
