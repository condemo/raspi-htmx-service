package handlers

import (
	"fmt"
	"net/http"

	"github.com/condemo/raspi-htmx-service/public/views/core"
)

type AuthHandler struct{}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) RegisterRoutes(r *http.ServeMux) {
	r.HandleFunc("GET /login", MakeHandler(h.loginPage))
	r.HandleFunc("POST /login", MakeHandler(h.login))
}

func (h *AuthHandler) loginPage(w http.ResponseWriter, r *http.Request) error {
	RenderTempl(w, r, core.Login())
	return nil
}

func (h *AuthHandler) login(w http.ResponseWriter, r *http.Request) error {
	un := r.FormValue("username")
	ps := r.FormValue("password")

	fmt.Println(un, ps)
	return nil
}
