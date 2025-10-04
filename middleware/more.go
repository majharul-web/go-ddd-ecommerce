package middleware

import (
	"log"
	"net/http"
)

func MoreMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("MoreMiddleware: Before handler")
		next.ServeHTTP(w, r)
		
	})
}
