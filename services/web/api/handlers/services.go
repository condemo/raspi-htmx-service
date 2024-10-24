package handlers

import (
	"fmt"
	"log"
	"net/http"

	manager "github.com/condemo/raspi-htmx-service/services/common/genproto/services"
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
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		res, err := h.managerConn.GetServices(r.Context(), &manager.GetServicesRequest{})
		if err != nil {
			log.Fatal("error in get services", err)
		}

		fmt.Fprint(w, res)
	})
}
