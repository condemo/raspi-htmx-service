package handlers

import (
	"net/http"

	"github.com/condemo/raspi-htmx-service/public/views/core"
)

type ViewHandler struct{}

func NewViewHandler() *ViewHandler {
	return &ViewHandler{}
}

func (h *ViewHandler) RegisterRoutes(r *http.ServeMux) {
	r.HandleFunc("/", h.homeView)
}

func (h *ViewHandler) homeView(w http.ResponseWriter, r *http.Request) {
	RenderTempl(w, r, core.Home())
}
