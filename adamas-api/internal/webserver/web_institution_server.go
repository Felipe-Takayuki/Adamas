package webserver

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/entity"
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/service"
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/utils"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
)

type WebInstitutionHandler struct {
	institutionService *service.InstitutionService
}

func NewWebInstiHandler(institutionService *service.InstitutionService) *WebInstitutionHandler {
	return &WebInstitutionHandler{
		institutionService: institutionService,
	}
}

func (wih *WebInstitutionHandler) CreateInstitution(w http.ResponseWriter, r *http.Request, tokenAuth *jwtauth.JWTAuth) {
	var institution *entity.Institution
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&institution)
	if err != nil {
		error := utils.ErrorMessage{Message: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
	result, err := wih.institutionService.CreateInstitution(institution.Name, institution.Email, institution.Password, institution.CNPJ)
	if err != nil {
		error := utils.ErrorMessage{Message: err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error)
		return
	} else {
		claims := map[string]interface{}{"id": result.ID, "name": result.Name, "email": result.Email, "user_type": result.UserType, "exp": jwtauth.ExpireIn(time.Minute * 10)}
		_, tokenString, _ = tokenAuth.Encode(claims)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"token": tokenString,
		})
	}

}

func (wih *WebInstitutionHandler) LoginInstitution(w http.ResponseWriter, r *http.Request, tokenAuth *jwtauth.JWTAuth) {
	var login *entity.Institution
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		error := utils.ErrorMessage{Message: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
	result, err := wih.institutionService.LoginInstitution(login.Email, login.Password)
	if err != nil {
		error := utils.ErrorMessage{Message: err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error)
		return
	} else {
		claims := map[string]interface{}{"id": result.ID, "name": result.Name, "email": result.Email, "user_type": result.UserType, "exp": jwtauth.ExpireIn(time.Minute * 10)}
		_, tokenString, _ = tokenAuth.Encode(claims)
		json.NewEncoder(w).Encode(map[string]interface{}{
			"token": tokenString,
		})
	}
}

func (wih *WebInstitutionHandler) GetInstitutionByID(w http.ResponseWriter, r *http.Request) {
	institutionID, err := strconv.Atoi(chi.URLParam(r, "institution_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	institution, err := wih.institutionService.GetInstitutionByID(int64(institutionID))
	if err != nil {
		error := utils.ErrorMessage{Message: err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error)
		return
	}
	json.NewEncoder(w).Encode(institution)
}
