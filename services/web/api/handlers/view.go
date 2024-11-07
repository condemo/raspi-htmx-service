package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/condemo/raspi-htmx-service/services/common/config"
	manager "github.com/condemo/raspi-htmx-service/services/common/genproto/services"
	sysinfo "github.com/condemo/raspi-htmx-service/services/common/genproto/services/sys_info"
	"github.com/condemo/raspi-htmx-service/services/web/public/views/core"
	"google.golang.org/grpc"
)

type ViewHandler struct {
	infoConn    sysinfo.SysInfoServiceClient
	managerConn manager.ServiceManagerClient
}

func NewViewHandler(ic *grpc.ClientConn, mc *grpc.ClientConn) *ViewHandler {
	inConn := sysinfo.NewSysInfoServiceClient(ic)
	mConn := manager.NewServiceManagerClient(mc)
	return &ViewHandler{
		infoConn:    inConn,
		managerConn: mConn,
	}
}

func (h *ViewHandler) RegisterRoutes(r *http.ServeMux) {
	r.HandleFunc("GET /", MakeHandler(h.homeView))
	r.HandleFunc("GET /config", MakeHandler(h.configView))
}

func (h *ViewHandler) homeView(w http.ResponseWriter, r *http.Request) error {
	si, err := h.infoConn.GetInfo(context.Background(), &sysinfo.GetInfoRequest{})
	if err != nil {
		log.Fatal("error getting info from `GetInfo` \n", err)
	}

	sl, err := h.managerConn.GetServices(r.Context(), &manager.GetServicesRequest{})
	if err != nil {
		return err
	}

	RenderTempl(w, r, core.Home(si.GetSisInfo(), sl.Services))
	return nil
}

func (h *ViewHandler) configView(w http.ResponseWriter, r *http.Request) error {
	RenderTempl(w, r, core.ConfigPage(config.UsConf))
	return nil
}
