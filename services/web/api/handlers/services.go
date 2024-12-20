package handlers

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/condemo/raspi-htmx-service/services/common/genproto/pb"
	"github.com/condemo/raspi-htmx-service/services/web/public/views/components"
	"google.golang.org/grpc"
)

type ServiceHandler struct {
	managerConn pb.ServiceManagerClient
}

func NewServiceHandler(mConn *grpc.ClientConn) *ServiceHandler {
	mc := pb.NewServiceManagerClient(mConn)
	return &ServiceHandler{
		managerConn: mc,
	}
}

func (h *ServiceHandler) RegisterRoutes(r *http.ServeMux) {
	r.HandleFunc("POST /start/{id}", MakeHandler(h.startService))
	r.HandleFunc("POST /stop/{id}", MakeHandler(h.stopService))
	r.HandleFunc("GET /full/{id}", MakeHandler(h.getFullInfo))
	r.HandleFunc("GET /config/{id}", MakeHandler(h.getConfig))
	r.HandleFunc("PUT /config/{id}", MakeHandler(h.updateConfig))
}

func (h *ServiceHandler) startService(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.ParseUint(r.PathValue("id"), 10, 32)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
	defer cancel()

	serv, err := h.managerConn.StartService(ctx, &pb.ServiceIdRequest{Id: uint32(id)})
	if err != nil {
		return err
	}

	return RenderTempl(w, r, components.ServiceCard(serv))
}

func (h *ServiceHandler) stopService(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.ParseUint(r.PathValue("id"), 10, 32)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
	defer cancel()

	serv, err := h.managerConn.StopService(ctx, &pb.ServiceIdRequest{Id: uint32(id)})
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

	res, err := h.managerConn.GetFullInfo(ctx, &pb.ServiceIdRequest{Id: uint32(id)})
	if err != nil {
		return err
	}

	return RenderTempl(w, r, components.FullInfoView(res))
}

func (h *ServiceHandler) getConfig(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 32)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
	defer cancel()

	res, err := h.managerConn.GetConfig(ctx, &pb.ServiceIdRequest{Id: uint32(id)})
	if err != nil {
		return err
	}

	return RenderTempl(w, r, components.ServiceConfigModal(res.GetConfig()))
}

func (h *ServiceHandler) updateConfig(w http.ResponseWriter, r *http.Request) error {
	city := r.FormValue("city")
	id, err := strconv.ParseInt(r.PathValue("id"), 10, 32)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
	defer cancel()

	_, err = h.managerConn.UpdateConfig(ctx, &pb.ServiceConfig{
		Id:   uint32(id),
		City: city,
	})

	return err
}
