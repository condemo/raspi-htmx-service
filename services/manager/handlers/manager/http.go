package handlers

import (
	"fmt"
	"net/http"

	manager "github.com/condemo/raspi-htmx-service/services/common/genproto/services"
	"github.com/condemo/raspi-htmx-service/services/manager/types"
)

type ManagerHttpHandler struct {
	serviceManager types.ServiceManager
}

func NewManagerHttpHandler(sm types.ServiceManager) *ManagerHttpHandler {
	return &ManagerHttpHandler{serviceManager: sm}
}

func (h *ManagerHttpHandler) RegisterRoutes(r *http.ServeMux) {
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "working")
	})
	r.HandleFunc("GET /register", h.RegisterServices)
	r.HandleFunc("GET /get-services", h.GetServices)
}

// TODO: Crear los dos handlers que hay implementados en la interfaz
func (h *ManagerHttpHandler) RegisterServices(w http.ResponseWriter, r *http.Request) {
	// TODO: Crear la struct desde la request
	req := &manager.RegisterServiceRequest{
		Id:   22,
		Name: "Servicio a ver",
	}

	if err := h.serviceManager.RegisterService(r.Context(), req); err != nil {
		fmt.Fprint(w, err)
		return
	}

	fmt.Fprint(w, "success")
}

func (h *ManagerHttpHandler) GetServices(w http.ResponseWriter, r *http.Request) {
	raspiServiceList := h.serviceManager.GetServices(r.Context(), &manager.GetServicesRequest{})

	fmt.Fprint(w, raspiServiceList)
}
