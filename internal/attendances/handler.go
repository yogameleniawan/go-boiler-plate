package attendances

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/base-go/backend/internal/shared/models"
	"github.com/base-go/backend/pkg/response"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return Handler{
		service: service,
	}
}

func (h Handler) CreateAttendance(w http.ResponseWriter, r *http.Request) {
	var attendance models.Attendance
	if err := json.NewDecoder(r.Body).Decode(&attendance); err != nil {
		response.ResponseError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	createdAttendance, statusCode, err := h.service.Create(r.Context(), attendance)
	if err != nil {
		response.ResponseError(w, statusCode, err.Error())
		return
	}

	response.ResponseJSON(w, statusCode, createdAttendance)
}

func (h Handler) GetAttendanceList(w http.ResponseWriter, r *http.Request) {
	data, statusCode, err := h.service.GetAll(r.Context())
	if err != nil {
		response.ResponseError(w, statusCode, err.Error())
		return
	}
	response.ResponseJSON(w, statusCode, data)
}

func (h Handler) GetAttendanceByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, "Invalid attendance ID")
		return
	}

	data, statusCode, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		response.ResponseError(w, statusCode, err.Error())
		return
	}

	response.ResponseJSON(w, statusCode, data)
}

func (h Handler) UpdateAttendance(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, "Invalid attendance ID")
		return
	}

	var attendance models.Attendance
	if err := json.NewDecoder(r.Body).Decode(&attendance); err != nil {
		response.ResponseError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	updatedAttendance, statusCode, err := h.service.Update(r.Context(), id, attendance)
	if err != nil {
		response.ResponseError(w, statusCode, err.Error())
		return
	}

	response.ResponseJSON(w, statusCode, updatedAttendance)
}

func (h Handler) DeleteAttendance(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.ResponseError(w, http.StatusBadRequest, "Invalid attendance ID")
		return
	}

	statusCode, err := h.service.Delete(r.Context(), id)
	if err != nil {
		response.ResponseError(w, statusCode, err.Error())
		return
	}

	response.ResponseJSON(w, statusCode, map[string]string{"message": "Attendance deleted successfully"})
}
