package middlewares

import (
	"net/http"
)

func Preflight(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// handle preflight requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(200)
			return
		}

		next.ServeHTTP(w, r)

	})
}
