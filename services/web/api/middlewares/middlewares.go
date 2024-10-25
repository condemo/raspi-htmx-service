package middlewares

import (
	"context"
	"net/http"

	"github.com/condemo/raspi-htmx-service/services/web/api/utils"
)

type wrapperResponse struct {
	http.ResponseWriter
	status int
}

func (wr *wrapperResponse) WriteHeader(status int) {
	wr.ResponseWriter.WriteHeader(status)
	wr.status = status
}

type Middleware func(next http.Handler) http.HandlerFunc

func MiddlewareStack(m ...Middleware) Middleware {
	return func(next http.Handler) http.HandlerFunc {
		for i := len(m) - 1; i >= 0; i-- {
			next = m[i](next)
		}
		return next.ServeHTTP
	}
}

func RequireAuth(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("raspi-token")
		if err != nil {
			http.Redirect(w, r, "/auth/login", http.StatusPermanentRedirect)
			return
		}

		if c.Value == "" {
			http.Redirect(w, r, "/auth/login", http.StatusPermanentRedirect)
			return
		}

		claims, err := utils.ValidateJWT(c.Value)
		if err != nil {
			http.Redirect(w, r, "/auth/login", http.StatusPermanentRedirect)
			return
		}

		id := claims.UserID
		ctx := context.WithValue(r.Context(), utils.UserID("userID"), id)

		next.ServeHTTP(w, r.Clone(ctx))
	}
}
