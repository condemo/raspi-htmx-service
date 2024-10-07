package handlers

import (
	"net/http"
	"strconv"

	"github.com/condemo/raspi-htmx-service/public/views/components"
	"github.com/condemo/raspi-htmx-service/services"
)

type ServiceHandler struct{}

func NewServiceHandler() *ServiceHandler {
	return &ServiceHandler{}
}

func (h *ServiceHandler) RegisterRoutes(r *http.ServeMux) {
	r.HandleFunc("POST /", MakeHandler(h.switchService))
}

func (h *ServiceHandler) switchService(w http.ResponseWriter, r *http.Request) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}

	st, err := strconv.ParseBool(r.FormValue("status"))
	if err != nil {
		return err
	}

	s, err := services.UpdateService(r.FormValue("name"), st)
	if err != nil {
		return err
	}

	return RenderTempl(w, r, components.ServiceCard(s))
}
