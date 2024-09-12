package middlewares

import "net/http"

type Middleware func(next http.Handler) http.HandlerFunc

func MiddlewareStack(m ...Middleware) Middleware {
	return func(next http.Handler) http.HandlerFunc {
		for i := len(m) - 1; i >= 0; i-- {
			next = m[i](next)
		}
		return next.ServeHTTP
	}
}
