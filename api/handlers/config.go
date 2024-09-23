package handlers

import (
	"net/http"

	"github.com/condemo/raspi-htmx-service/config"
	"github.com/condemo/raspi-htmx-service/public/views/components"
)

type ConfigHandler struct{}

func NewConfigHandler() *ConfigHandler {
	return &ConfigHandler{}
}

func (h *ConfigHandler) RegisterRoutes(r *http.ServeMux) {
	// TODO: Replantear toda la vaina desde 0
	r.HandleFunc("GET /", MakeHandler(h.getConfig))
	r.HandleFunc("PUT /", MakeHandler(h.updateConfig))
}

func (h *ConfigHandler) getConfig(w http.ResponseWriter, r *http.Request) error {
	RenderTempl(w, r, components.ConfigSection())
	return nil
}

func (h *ConfigHandler) updateConfig(w http.ResponseWriter, r *http.Request) error {
	config.SaveConf()
	return nil
}
