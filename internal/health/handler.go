package health

import (
	"net/http"

	"github.com/absendulu-project/backend/pkg/response"
)

type Handler interface {
	Health(w http.ResponseWriter, r *http.Request)
}

type hander struct {
	Service Service
}

func NewHandler(service Service) Handler {
	return &hander{
		service,
	}
}

func (h *hander) Health(w http.ResponseWriter, r *http.Request) {

	response.ResponseJSON(w, http.StatusOK, response.JSON{
		Status:  true,
		Message: "Success Get Health",
		Data:    nil,
	})
}
