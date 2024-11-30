package handlers

import (
	"context"
	"net/http"
	"strconv"
	"time"

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
	r.HandleFunc("GET /full/{id}", MakeHandler(h.getFullInfo))
}

func (h *ServiceHandler) startService(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 32)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
	defer cancel()

	serv, err := h.managerConn.StartService(ctx, &manager.ServiceIdRequest{Id: int32(id)})
	if err != nil {
		return err
	}

	return RenderTempl(w, r, components.ServiceCard(serv))
}

func (h *ServiceHandler) stopService(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 32)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
	defer cancel()

	serv, err := h.managerConn.StopService(ctx, &manager.ServiceIdRequest{Id: int32(id)})
	if err != nil {
		return err
	}

	return RenderTempl(w, r, components.ServiceCard(serv))
}

func (h *ServiceHandler) getFullInfo(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 32)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
	defer cancel()

	res, err := h.managerConn.GetFullInfo(ctx, &manager.ServiceIdRequest{Id: int32(id)})
	if err != nil {
		return err
	}

	return RenderTempl(w, r, components.FullInfoView(res))
}
