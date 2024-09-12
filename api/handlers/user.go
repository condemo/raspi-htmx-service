package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/condemo/raspi-htmx-service/api/handlers/errors"
	"github.com/condemo/raspi-htmx-service/api/utils"
	"github.com/condemo/raspi-htmx-service/public/views/core"
	"github.com/condemo/raspi-htmx-service/store"
	"github.com/condemo/raspi-htmx-service/types"
)

type AuthHandler struct {
	store store.Store
}

func NewAuthHandler(s store.Store) *AuthHandler {
	return &AuthHandler{store: s}
}

func (h *AuthHandler) RegisterRoutes(r *http.ServeMux) {
	r.HandleFunc("GET /login", MakeHandler(h.loginPage))
	r.HandleFunc("POST /login", MakeHandler(h.login))
	r.HandleFunc("POST /signup", MakeHandler(h.signup))
	r.HandleFunc("POST /logout", MakeHandler(h.logout))
}

func (h *AuthHandler) loginPage(w http.ResponseWriter, r *http.Request) error {
	RenderTempl(w, r, core.Login())
	return nil
}

func (h *AuthHandler) login(w http.ResponseWriter, r *http.Request) error {
	un := r.FormValue("username")
	ps := r.FormValue("password")

	user, err := h.store.GetUserByUsername(un)
	if err != nil {
		return errors.NewUINotifyError("user not found", http.StatusNotFound)
	}

	if !utils.VerifyPass(user.Password, ps) {
		return errors.NewUINotifyError("invalid password", http.StatusUnauthorized)
	}

	token, expTime, err := utils.CreateJWT(user.ID)
	if err != nil {
		return err
	}

	c := http.Cookie{
		Name:    "raspi-token",
		Value:   token,
		Path:    "/",
		Expires: expTime,
	}

	http.SetCookie(w, &c)
	w.Header().Set("HX-Redirect", "/app")
	return nil
}

func (h *AuthHandler) signup(w http.ResponseWriter, r *http.Request) error {
	un := r.FormValue("username")
	pass, err := utils.EncryptPass(r.FormValue("password"))

	fmt.Println(un)
	fmt.Println(pass)
	if err != nil {
		return err
	}

	user := &types.User{
		Username: un,
		Password: pass,
	}

	err = h.store.CreateUser(user)
	if err != nil {
		return err
	}

	w.Header().Add("HX-Redirect", "/auth/login")
	w.WriteHeader(http.StatusCreated)

	return nil
}

func (h *AuthHandler) logout(w http.ResponseWriter, r *http.Request) error {
	c := http.Cookie{
		Name:    "raspi-token",
		Value:   "",
		Path:    "/",
		Expires: time.Unix(0, 0),
	}

	http.SetCookie(w, &c)
	w.Header().Set("HX-Redirect", "/auth/login")
	w.WriteHeader(http.StatusOK)
	return nil
}
