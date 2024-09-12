package handlers

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
	"github.com/condemo/raspi-htmx-service/api/handlers/errors"
	"github.com/condemo/raspi-htmx-service/public/views/core"
)

type CustomHandler func(w http.ResponseWriter, r *http.Request) error

func MakeHandler(f CustomHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			switch err := err.(type) {
			case errors.UINofifyError:
				w.WriteHeader(err.Status)
			case errors.ApiError:
				RenderTempl(w, r, core.ErrorPage(err))
			default:
				fmt.Fprintf(w, "unformat error -> %s", err.Error())
			}
			slog.Error("API ERROR", "err", err.Error(), "path", r.URL.Path)
		}
	}
}

func RenderTempl(w http.ResponseWriter, r *http.Request, c templ.Component) {
	c.Render(r.Context(), w)
}

func TextResponse(w http.ResponseWriter, msg string, status int) {
	w.WriteHeader(status)
	fmt.Fprint(w, msg)
}
