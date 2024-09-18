package middlewares

import (
	"fmt"
	"net/http"
)

func Recover(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)

				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(w, "internal server error")
			}
		}()
		next.ServeHTTP(w, r)
	}
}
