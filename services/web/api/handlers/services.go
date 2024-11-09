package handlers

import (
	"net/http"
	"strconv"

	manager "github.com/condemo/raspi-htmx-service/services/common/genproto/services"
	"github.com/condemo/raspi-htmx-service/services/web/public/views/components"
	"google.golang.org/grpc"
)

type ServiceHandler struct {
	managerConn manager.ServiceManagerClient
}

func NewServiceHandler(mConn *grpc.ClientConn) *ServiceHandler {
	mc := manager.NewServiceManagerClient(mConn)
	return &ServiceHandler{
		managerConn: mc,
	}
}

func (h *ServiceHandler) RegisterRoutes(r *http.ServeMux) {
	r.HandleFunc("POST /start/{id}", MakeHandler(h.startService))
	r.HandleFunc("POST /stop/{id}", MakeHandler(h.stopService))
}

func (h *ServiceHandler) startService(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 32)
	if err != nil {
		return err
	}

	serv, err := h.managerConn.StartService(r.Context(), &manager.ServiceIdRequest{Id: int32(id)})
	if err != nil {
		return err
	}

	// TODO: Mockup, conseguir datos reales, quizás startService y stopService deberían
	// deverían devolver el RaspiService otra vez y renderizar la tarjeta entera

	return RenderTempl(w, r, components.ServiceCard(serv))
}

func (h *ServiceHandler) stopService(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 32)
	if err != nil {
		return err
	}

	serv, err := h.managerConn.StopService(r.Context(), &manager.ServiceIdRequest{Id: int32(id)})
	if err != nil {
		return err
	}

	// TODO: Mockup, conseguir datos reales, quizás startService y stopService deberían
	// deverían devolver el RaspiService otra vez y renderizar la tarjeta entera
	return RenderTempl(w, r, components.ServiceCard(serv))
}
