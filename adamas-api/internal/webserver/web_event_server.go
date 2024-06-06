package webserver

import (
	"encoding/json"
	"net/http"

	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/entity/reqs"
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/service"
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/utils"
	"github.com/go-chi/jwtauth"
)

type WebEventHandler struct {
	eventService *service.EventService
}

func NewWebEventHandler(eventService *service.EventService) *WebEventHandler {
	return &WebEventHandler{
		eventService: eventService,
	}
}

func (weh *WebEventHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	flt64, ok := claims["id"].(float64)
	if !ok {
		http.Error(w, "id is not int!", http.StatusInternalServerError)
		return
	}
	var req *reqs.CreateEventRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		error := utils.ErrorMessage{Message: err.Error()}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
	result, err := weh.eventService.CreateEvent(req.Name, req.Address, req.Date, req.Description, int(flt64))
	if err != nil {
		error := utils.ErrorMessage{Message: err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error)
		return
	}
	json.NewEncoder(w).Encode(result)
}
