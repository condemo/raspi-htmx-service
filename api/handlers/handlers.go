package handlers

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
	"github.com/condemo/raspi-htmx-service/api/handlers/errors"
	"github.com/condemo/raspi-htmx-service/public/views/components"
	"github.com/condemo/raspi-htmx-service/public/views/core"
)

type CustomHandler func(w http.ResponseWriter, r *http.Request) error

func MakeHandler(f CustomHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			switch err := err.(type) {
			case errors.UINofifyError:
				// Si el status code no es positivo htmx lo trata como un error
				// w.WriteHeader(err.Status)
				RenderTempl(w, r, components.SimpleError(err))
			case errors.ApiError:
				RenderTempl(w, r, components.ApiError(err))
			default:
				// FIX: No enviar los errores tal cual al server
				RenderTempl(w, r, core.ErrorPage(errors.NewUnformatError(http.StatusInternalServerError)))
				slog.Error("Interal", "unformat", err.Error(), "path", r.URL.Path)
			}
			slog.Error("API ERROR", "err", err.Error(), "path", r.URL.Path)
		}
	}
}

func RenderTempl(w http.ResponseWriter, r *http.Request, c templ.Component) error {
	c.Render(r.Context(), w)
	return nil
}

func TextResponse(w http.ResponseWriter, msg string, status int) {
	w.WriteHeader(status)
	fmt.Fprint(w, msg)
}
