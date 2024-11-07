package handlers

import (
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
	// TODO: Definir las rutas para comunicarse con el `ManagerService`
}
