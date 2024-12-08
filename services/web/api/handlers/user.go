package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/condemo/raspi-htmx-service/services/common/genproto/pb"
	"github.com/condemo/raspi-htmx-service/services/common/store"
	"github.com/condemo/raspi-htmx-service/services/web/api/handlers/errors"
	"github.com/condemo/raspi-htmx-service/services/web/api/utils"
	"github.com/condemo/raspi-htmx-service/services/web/public/views/core"
	"github.com/condemo/raspi-htmx-service/services/web/types"
	"google.golang.org/grpc"
)

type AuthHandler struct {
	store   store.Store
	logConn pb.LoggerServiceClient
}

func NewAuthHandler(s store.Store, logC *grpc.ClientConn) *AuthHandler {
	lc := pb.NewLoggerServiceClient(logC)
	return &AuthHandler{
		store:   s,
		logConn: lc,
	}
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
		m := fmt.Sprintf("user not found error - %s", err)
		h.logConn.LogMessage(r.Context(), utils.MakeLog(
			pb.LogMessageType_ERROR, m))
		return errors.NewUINotifyError("user not found", http.StatusNotFound)
	}

	if !utils.VerifyPass(user.Password, ps) {
		m := fmt.Sprintf("invalid password - %s", err)
		h.logConn.LogMessage(r.Context(), utils.MakeLog(
			pb.LogMessageType_ERROR, m))
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

	w.Header().Add("HX-Redirect", "/app/")
	w.WriteHeader(http.StatusOK)

	m := fmt.Sprintf("User `%s` Succesfull Logged", user.Username)
	h.logConn.LogMessage(r.Context(), utils.MakeLog(
		pb.LogMessageType_SUCCESS, m))
	return nil
}

func (h *AuthHandler) signup(w http.ResponseWriter, r *http.Request) error {
	un := r.FormValue("username")
	pass, err := utils.EncryptPass(r.FormValue("password"))
	if err != nil {
		h.logConn.LogMessage(r.Context(), utils.MakeLog(
			pb.LogMessageType_ERROR, err.Error()))
		return err
	}

	user := &types.User{
		Username: un,
		Password: pass,
	}

	err = h.store.CreateUser(user)
	if err != nil {
		h.logConn.LogMessage(r.Context(), utils.MakeLog(
			pb.LogMessageType_ERROR, err.Error()))
		return err
	}

	w.Header().Add("HX-Redirect", "/auth/login")
	w.WriteHeader(http.StatusCreated)

	m := fmt.Sprintf("User `%s` Succesfull Created", user.Username)
	h.logConn.LogMessage(r.Context(), utils.MakeLog(
		pb.LogMessageType_SUCCESS, m))
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

	h.logConn.LogMessage(r.Context(), utils.MakeLog(
		pb.LogMessageType_INFO, "User Succesfull Logout"))
	return nil
}
