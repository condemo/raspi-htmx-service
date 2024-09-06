package handlers

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
)

func RenderTempl(w http.ResponseWriter, r *http.Request, c templ.Component) {
	c.Render(r.Context(), w)
}

func TextResponse(w http.ResponseWriter, msg string, status int) {
	w.WriteHeader(status)
	fmt.Fprint(w, msg)
}
