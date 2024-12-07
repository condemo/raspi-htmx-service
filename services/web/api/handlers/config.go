package handlers

import (
	"net/http"

	"github.com/condemo/raspi-htmx-service/services/common/config"
	"github.com/condemo/raspi-htmx-service/services/common/genproto/pb"
	"github.com/condemo/raspi-htmx-service/services/web/api/utils"
	"github.com/condemo/raspi-htmx-service/services/web/public/views/components"
	"google.golang.org/grpc"
)

type ConfigHandler struct {
	logConn pb.LoggerServiceClient
}

func NewConfigHandler(logC *grpc.ClientConn) *ConfigHandler {
	lc := pb.NewLoggerServiceClient(logC)
	return &ConfigHandler{
		logConn: lc,
	}
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
		h.logConn.LogMessage(r.Context(), utils.MakeLog(
			pb.LogMessageType_ERROR, err.Error()))
		return err
	}

	if err := config.UpdateConf(*us); err != nil {
		h.logConn.LogMessage(r.Context(), utils.MakeLog(
			pb.LogMessageType_ERROR, err.Error()))
		return err
	}

	w.Header().Set("HX-Redirect", "/app/config")
	w.WriteHeader(http.StatusAccepted)
	h.logConn.LogMessage(r.Context(), utils.MakeLog(
		pb.LogMessageType_SUCCESS, "Config Successfull Updated"))
	return nil
}
