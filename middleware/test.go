package middleware

import (
	"log"
	"net/http"
)

func TestingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("TestingMiddleware: Before handler")
		next.ServeHTTP(w, r)
		
	})
}
