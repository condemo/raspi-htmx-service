package handlers

import (
	"context"
	"log"

	manager "github.com/condemo/raspi-htmx-service/services/common/genproto/services"
	"github.com/condemo/raspi-htmx-service/services/common/genproto/services/logger"
	raspiservices "github.com/condemo/raspi-htmx-service/services/common/genproto/services/raspi_services"
	"github.com/condemo/raspi-htmx-service/services/common/util"
	"github.com/condemo/raspi-htmx-service/services/manager/logs"
	"github.com/condemo/raspi-htmx-service/services/manager/types"
	"google.golang.org/grpc"
)

// PERF: Mejorar toda la implementación del handler y del servico/interface, es un lío
type ManagerGrpcHandler struct {
	manager.UnimplementedServiceManagerServer
	serviceManager types.ServiceManager
	weatherService raspiservices.WeatherServiceClient
	logService     logger.LoggerServiceClient
}

func NewManagerGrpcHandler(grpc *grpc.Server, sm types.ServiceManager) {
	// Load all the conns
	weatherGrpc := util.NewGrpcClient(":8010")
	weatherConn := raspiservices.NewWeatherServiceClient(weatherGrpc)

	logGrpc := util.NewGrpcClient(":7000")
	logConn := logger.NewLoggerServiceClient(logGrpc)

	gRPCHandler := &ManagerGrpcHandler{
		serviceManager: sm,
		weatherService: weatherConn,
		logService:     logConn,
	}

	_, err := gRPCHandler.logService.LogMessage(context.Background(), logs.MakeLog(
		logger.MessageType_SUCCESS, "Manager Handler Starts"))
	if err != nil {
		log.Fatal("error in logger", err)
	}

	// TODO: Load/Read all `RaspiServices` - Cutre
	if err := gRPCHandler.LoadServices(context.Background()); err != nil {
		log.Fatal("error loading services in manager - ", err)
	}

	manager.RegisterServiceManagerServer(grpc, gRPCHandler)
}

// PERF: `LoadServices` no debería estar dentro del handler sino que debaria llamarse
// antes de crear `ManagerGrpcHandler`
func (h *ManagerGrpcHandler) LoadServices(ctx context.Context) error {
	ws, err := h.weatherService.GetStatus(ctx, &raspiservices.EmptyRequest{})
	if err != nil {
		return err
	}

	// TODO: Cambiar `LoadService` por `LoadServices`; ir obteniendo la data de cada servicio
	// para luego cargarlo de una con una sola llamada a la función
	h.serviceManager.LoadService(ctx, &manager.RaspiService{
		Id: ws.Id, Status: ws.Status, Name: ws.Name,
		Data: &manager.ServiceCardData{
			Icon:        ws.GetData().GetIcon(),
			DataText:    ws.GetData().GetData(),
			LastUpdated: ws.GetData().GetLastUpdated(),
		},
	})

	return nil
}

func (h *ManagerGrpcHandler) GetServices(ctx context.Context, req *manager.GetServicesRequest) (*manager.GetServicesResponse, error) {
	err := h.LoadServices(ctx)
	if err != nil {
		return nil, err
	}

	sl := h.serviceManager.GetServices(ctx)
	res := &manager.GetServicesResponse{Services: sl}
	return res, nil
}

func (h *ManagerGrpcHandler) StartService(ctx context.Context, req *manager.ServiceIdRequest) (*manager.RaspiService, error) {
	st, err := h.weatherService.Start(ctx, &raspiservices.EmptyRequest{})
	res := &manager.RaspiService{
		Id: st.Id, Name: st.Name, Status: st.Status,
		Data: &manager.ServiceCardData{
			Icon:        st.Data.Icon,
			DataText:    st.Data.Data,
			LastUpdated: st.GetData().GetLastUpdated(),
		},
	}
	return res, err
}

func (h *ManagerGrpcHandler) StopService(ctx context.Context, req *manager.ServiceIdRequest) (*manager.RaspiService, error) {
	st, err := h.weatherService.Stop(ctx, &raspiservices.EmptyRequest{})
	res := &manager.RaspiService{
		Id: st.Id, Name: st.Name, Status: st.Status,
		Data: &manager.ServiceCardData{
			Icon:        st.Data.Icon,
			DataText:    st.Data.Data,
			LastUpdated: st.GetData().GetLastUpdated(),
		},
	}

	return res, err
}
