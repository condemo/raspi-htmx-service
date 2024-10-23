package handlers

import (
	"net/http"

	"github.com/condemo/raspi-htmx-service/services/web/api/utils"
	"github.com/condemo/raspi-htmx-service/services/web/config"
	"github.com/condemo/raspi-htmx-service/services/web/public/views/components"
)

type ConfigHandler struct{}

func NewConfigHandler() *ConfigHandler {
	return &ConfigHandler{}
}

func (h *ConfigHandler) RegisterRoutes(r *http.ServeMux) {
	r.HandleFunc("GET /", MakeHandler(h.getConfig))
	r.HandleFunc("PUT /", MakeHandler(h.updateConfig))
}

func (h *ConfigHandler) getConfig(w http.ResponseWriter, r *http.Request) error {
	RenderTempl(w, r, components.ConfigSection(config.UsConf))
	return nil
}

func (h *ConfigHandler) updateConfig(w http.ResponseWriter, r *http.Request) error {
	us := &config.UserConfig{}
	err := utils.ConfigParser(r, us)
	if err != nil {
		return err
	}

	if err := config.UpdateConf(*us); err != nil {
		return err
	}

	w.Header().Set("HX-Redirect", "/app/config")
	w.WriteHeader(http.StatusAccepted)

	return nil
}
