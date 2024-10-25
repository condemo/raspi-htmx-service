package handlers

import (
	"net/http"

	"github.com/condemo/raspi-htmx-service/services/common/config"
	"github.com/condemo/raspi-htmx-service/services/web/public/views/core"
)

type ViewHandler struct{}

func NewViewHandler() *ViewHandler {
	return &ViewHandler{}
}

func (h *ViewHandler) RegisterRoutes(r *http.ServeMux) {
	r.HandleFunc("GET /", MakeHandler(h.homeView))
	r.HandleFunc("GET /config", MakeHandler(h.configView))
}

func (h *ViewHandler) homeView(w http.ResponseWriter, r *http.Request) error {
	RenderTempl(w, r, core.Home())
	return nil
}

func (h *ViewHandler) configView(w http.ResponseWriter, r *http.Request) error {
	RenderTempl(w, r, core.ConfigPage(config.UsConf))
	return nil
}
