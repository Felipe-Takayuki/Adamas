package webserver

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/entity"
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/entity/reqs"
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/service"
	"github.com/Felipe-Takayuki/Adamas/adamas-api/internal/utils"
	"github.com/go-chi/chi"
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

func (weh *WebEventHandler) GetEventByName(w http.ResponseWriter, r *http.Request) {
	eventName := chi.URLParam(r, "event")
	if eventName == "" {
		error := utils.ErrorMessage{Message: "name event is required"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
	events, err := weh.eventService.GetEventByName(eventName)
	if err != nil {
		error := utils.ErrorMessage{Message: err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error)
		return
	}
	json.NewEncoder(w).Encode(events)
}

func (weh *WebEventHandler) GetEventByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	eventID, err := strconv.Atoi(chi.URLParam(r, "event_id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	event, err := weh.eventService.GetEventByID(int64(eventID))
	if err != nil {
		error := utils.ErrorMessage{Message: err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error)
		return
	}
	json.NewEncoder(w).Encode(event)

}
func (weh *WebEventHandler) GetEvents(w http.ResponseWriter, r *http.Request) {
	events, err := weh.eventService.EventDB.GetEvents()
	if err != nil {
		error := utils.ErrorMessage{Message: err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error)
		return
	}
	json.NewEncoder(w).Encode(events)
}

func (weh *WebEventHandler) GetSubscribers(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	w.Header().Set("Content-Type", "application/json")
	userType, ok := claims["user_type"].(string)
	if !ok {
		http.Error(w, "user_type is not exists!", http.StatusInternalServerError)
		return
	}
	if userType == "institution_user" {
		eventID, err := strconv.Atoi(chi.URLParam(r, "event_id"))
		if err != nil {
			http.Error(w, "event_id is not int!", http.StatusInternalServerError)
			return
		}
		institutionID, ok := claims["id"].(float64)
		if !ok {
			http.Error(w, "id is not int!", http.StatusInternalServerError)
			return
		}
		subscribers, err := weh.eventService.GetSubscribersByEventID(int64(eventID), int64(institutionID))
		if err != nil {
			error := utils.ErrorMessage{Message: err.Error()}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(error)
			return
		}
		json.NewEncoder(w).Encode(subscribers)
	} else {
		error := utils.ErrorMessage{Message: "este usuário não possui essa permissão!"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
}

func (weh *WebEventHandler) GetPendingProjectsInEvent(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	w.Header().Set("Content-Type", "application/json")
	userType, ok := claims["user_type"].(string)
	if !ok {
		http.Error(w, "id is not exists!", http.StatusInternalServerError)
		return
	}
	if userType == "institution_user" {
		eventID, err := strconv.Atoi(chi.URLParam(r, "event_id"))
		if err != nil {
			http.Error(w, "event_id is not int!", http.StatusInternalServerError)
			return
		}
		institutionID, ok := claims["id"].(float64)
		if !ok {
			http.Error(w, "id is not int!", http.StatusInternalServerError)
			return
		}
		pendingProjects, err := weh.eventService.GetPendingProjectsInEvent(int64(eventID),int64(institutionID))
		if err != nil {
			error := utils.ErrorMessage{Message: err.Error()}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(error)
			return
		}
		json.NewEncoder(w).Encode(pendingProjects)
	} else {
		error := utils.ErrorMessage{Message: "este usuário não possui essa permissão!"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
}

func (weh *WebEventHandler) GetProjectsInEvent(w http.ResponseWriter, r *http.Request) {
	eventID, err := strconv.Atoi(chi.URLParam(r, "event_id"))
		if err != nil {
			http.Error(w, "event_id is not int!", http.StatusInternalServerError)
			return
	}
	approvedProjects, err := weh.eventService.GetProjectsInEvent(int64(eventID))
	if err != nil {
		error := utils.ErrorMessage{Message: err.Error()}
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(error)
		return 
	}
	json.NewEncoder(w).Encode(approvedProjects)
}

func (weh *WebEventHandler) CreateEvent(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	w.Header().Set("Content-Type", "application/json")
	userType, ok := claims["user_type"].(string)
	if !ok {
		http.Error(w, "id is not exists!", http.StatusInternalServerError)
		return
	}
	if userType == "institution_user" {
		institutionID, ok := claims["id"].(float64)
		if !ok {
			http.Error(w, "id is not int!", http.StatusInternalServerError)
			return
		}
		var req *entity.Event
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			error := utils.ErrorMessage{Message: err.Error()}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(error)
			return
		}
		event, err := weh.eventService.CreateEvent(req.Name, req.Address, req.StartDate, req.EndDate, req.Description, int64(institutionID))
		if err != nil {
			error := utils.ErrorMessage{Message: err.Error()}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(error)
			return
		}
		json.NewEncoder(w).Encode(event)
	} else {
		error := utils.ErrorMessage{Message: "este usuário não possui essa permissão!"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
}

func (weh *WebEventHandler) AddRoomInEvent(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	w.Header().Set("Content-Type", "application/json")
	userType, ok := claims["user_type"].(string)
	if !ok {
		http.Error(w, "user_type is not exists!", http.StatusInternalServerError)
		return
	}
	if userType == "institution_user" {
		institutionID, ok := claims["id"].(float64)
		if !ok {
			http.Error(w, "id is not int!", http.StatusInternalServerError)
			return
		}
		eventID, err := strconv.Atoi(chi.URLParam(r, "event_id"))
		if err != nil {
			http.Error(w, "event_id is not int!", http.StatusInternalServerError)
			return
		}
		var req *reqs.AddRoomRequest
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			error := utils.ErrorMessage{Message: err.Error()}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(error)
			return
		}
		newRoom, err := weh.eventService.AddRoomInEvent(int64(eventID), int64(institutionID), req.Name, req.QuantityProjects)
		if err != nil {
			error := utils.ErrorMessage{Message: err.Error()}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(error)
			return
		}
		json.NewEncoder(w).Encode(newRoom)
	} else {
		error := utils.ErrorMessage{Message: "este usuário não possui essa permissão!"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
}

func (weh *WebEventHandler) GetRoomsByEventID(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	w.Header().Set("Content-Type", "application/json")
	userType, ok := claims["user_type"].(string)
	if !ok {
		http.Error(w, "user_type is not exists!", http.StatusInternalServerError)
		return
	}
	if userType == "institution_user" {
		eventID, err := strconv.Atoi(chi.URLParam(r, "event_id"))
		if err != nil {
			http.Error(w, "event_id is not int!", http.StatusInternalServerError)
			return
		}
		institutionID, ok := claims["id"].(float64)
		if !ok {
			http.Error(w, "id is not int!", http.StatusInternalServerError)
			return
		}
		rooms, err := weh.eventService.GetRoomsByEventID(int64(eventID), int64(institutionID))
		if err != nil {
			error := utils.ErrorMessage{Message: err.Error()}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(error)
			return 
		}
		json.NewEncoder(w).Encode(rooms)
	} else {
		error := utils.ErrorMessage{Message: "este usuário não possui essa permissão!"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
}

func (weh *WebEventHandler) DeleteRoom(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	w.Header().Set("Content-Type", "application/json")
	userType, ok := claims["user_type"].(string)
	if !ok {
		http.Error(w, "user_type is not exists!", http.StatusInternalServerError)
		return
	}
	if userType == "institution_user" {
		eventID, err := strconv.Atoi(chi.URLParam(r, "event_id"))
		if err != nil {
			http.Error(w, "event_id is not int!", http.StatusInternalServerError)
			return
		}
		var req *entity.RoomEvent
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			error := utils.ErrorMessage{Message: err.Error()}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(error)
			return
		}

		err = weh.eventService.DeleteRoom(int64(req.ID), int64(eventID))
		if err != nil {
			error := utils.ErrorMessage{Message: err.Error()}
			w.WriteHeader(http.StatusInternalServerError)
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

func (weh *WebEventHandler) EventRegistration(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	w.Header().Set("Content-Type", "application/json")
	userType, ok := claims["user_type"].(string)
	if !ok {
		http.Error(w, "id is not exists!", http.StatusInternalServerError)
		return
	}
	if userType == "common_user" {
		eventID, err := strconv.Atoi(chi.URLParam(r, "event_id"))
		if err != nil {
			http.Error(w, "event_id is not int!", http.StatusInternalServerError)
			return
		}
		userID, ok := claims["id"].(float64)
		if !ok {
			http.Error(w, "id is not int!", http.StatusInternalServerError)
			return
		}
		newRegister, err := weh.eventService.EventRegistration(int64(eventID), int64(userID))
		if err != nil {
			error := utils.ErrorMessage{Message: err.Error()}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(error)
			return
		}
		json.NewEncoder(w).Encode(newRegister)
	} else {
		error := utils.ErrorMessage{Message: "este usuário não possui essa permissão!"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
}

func (weh *WebEventHandler) EventRequestParticipation(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	w.Header().Set("Content-Type", "application/json")
	userType, ok := claims["user_type"].(string)
	if !ok {
		http.Error(w, "id is not exists!", http.StatusInternalServerError)
		return
	}

	if userType == "common_user" {
		eventID, err := strconv.Atoi(chi.URLParam(r, "event_id"))
		if err != nil {
			http.Error(w, "event_id is not int!", http.StatusInternalServerError)
			return
		}
		userID, ok := claims["id"].(float64)
		if !ok {
			http.Error(w, "id is not int!", http.StatusInternalServerError)
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
		newRequestParty, err := weh.eventService.EventRequestParticipation(int64(eventID), int64(userID), req.ID)
		if err != nil {
			error := utils.ErrorMessage{Message: err.Error()}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(error)
			return
		}
		json.NewEncoder(w).Encode(newRequestParty)
	} else {
		error := utils.ErrorMessage{Message: "este usuário não possui essa permissão!"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
}

func (weh *WebEventHandler) ApproveParticipation(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	w.Header().Set("Content-Type", "application/json")
	userType, ok := claims["user_type"].(string)
	if !ok {
		http.Error(w, "id is not exists!", http.StatusInternalServerError)
		return
	}
	if userType == "institution_user" {
		eventID, err := strconv.Atoi(chi.URLParam(r, "event_id"))
		if err != nil {
			http.Error(w, "event_id is not int!", http.StatusInternalServerError)
			return
		}
		var req *reqs.ApproveProjectRequest
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			error := utils.ErrorMessage{Message: err.Error()}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(error)
			return
		}
		userID, ok := claims["id"].(float64)
		if !ok {
			http.Error(w, "id is not int!", http.StatusInternalServerError)
			return
		}
		projectApproved, err := weh.eventService.ApproveParticipation(req.ProjectID, int64(userID), int64(eventID), req.RoomID)
		if err != nil {
			error := utils.ErrorMessage{Message: err.Error()}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(error)
			return
		}
		json.NewEncoder(w).Encode(projectApproved)
	} else {
		error := utils.ErrorMessage{Message: "este usuário não possui essa permissão!"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
}

func (weh *WebEventHandler) EditRoom(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	w.Header().Set("Content-Type", "application/json")
	userType, ok := claims["user_type"].(string)
	if !ok {
		http.Error(w, "id is not exists!", http.StatusInternalServerError)
		return
	}

	if userType == "institution_user" {
		eventID, err := strconv.Atoi(chi.URLParam(r, "event_id"))
		if err != nil {
			http.Error(w, "event_id is not int!", http.StatusInternalServerError)
			return
		}
		ownerID, ok := claims["id"].(float64)
		if !ok {
			http.Error(w, "id is not int!", http.StatusInternalServerError)
			return
		}
		var req *entity.RoomEvent
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			error := utils.ErrorMessage{Message: err.Error()}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(error)
			return
		}
		roomEdited, err := weh.eventService.EditRoom(req.ID, int64(eventID), int64(req.QuantityProjects), int64(ownerID), req.Name)
		if err != nil {
			error := utils.ErrorMessage{Message: err.Error()}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(error)
			return
		}
		json.NewEncoder(w).Encode(roomEdited)
	} else {
		error := utils.ErrorMessage{Message: "este usuário não possui essa permissão!"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
}

func (weh *WebEventHandler) EditEvent(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	w.Header().Set("Content-Type", "application/json")
	userType, ok := claims["user_type"].(string)
	if !ok {
		http.Error(w, "id is not exists!", http.StatusInternalServerError)
		return
	}
	if userType == "institution_user" {
		eventID, err := strconv.Atoi(chi.URLParam(r, "event_id"))
		if err != nil {
			http.Error(w, "event_id is not int!", http.StatusInternalServerError)
			return
		}
		ownerID, ok := claims["id"].(float64)
		if !ok {
			http.Error(w, "id is not int!", http.StatusInternalServerError)
			return
		}
		var req *entity.Event
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			error := utils.ErrorMessage{Message: err.Error()}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(error)
			return
		}
		eventEdited, err := weh.eventService.EditEvent(int64(eventID), int64(ownerID), req.Name, req.Address, req.StartDate, req.EndDate, req.Description)
		if err != nil {
			error := utils.ErrorMessage{Message: err.Error()}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(error)
			return
		}
		json.NewEncoder(w).Encode(eventEdited)
	} else {
		error := utils.ErrorMessage{Message: "este usuário não possui essa permissão!"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
}

func (weh *WebEventHandler) DeleteEvent(w http.ResponseWriter, r *http.Request) {
	_, claims, _ := jwtauth.FromContext(r.Context())
	w.Header().Set("Content-Type", "application/json")
	userType, ok := claims["user_type"].(string)
	if !ok {
		http.Error(w, "id is not exists!", http.StatusInternalServerError)
		return
	}
	if userType == "institution_user" {
		eventID, err := strconv.Atoi(chi.URLParam(r, "event_id"))
		if err != nil {
			http.Error(w, "event_id is not int!", http.StatusInternalServerError)
			return
		}

		var req *entity.Institution
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			error := utils.ErrorMessage{Message: err.Error()}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(error)
			return
		}
		err = weh.eventService.DeleteEvent(int64(eventID), req.Email, req.Password)
		if err != nil {
			error := utils.ErrorMessage{Message: err.Error()}
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(error)
			return
		}
		json.NewEncoder(w).Encode(nil)
	} else {
		error := utils.ErrorMessage{Message: "este usuário não possui essa permissão!"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(error)
		return
	}
}
