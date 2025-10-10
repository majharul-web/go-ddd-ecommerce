package middlewares

import (
	"crypto/hmac"
	"crypto/sha256"
	"ecommerce/util"
	"net/http"
	"strings"
)

func (m *Middlewares) AuthenticateJWT(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		header := r.Header.Get("Authorization")
		if header == "" {
			http.Error(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		headerParts := strings.Split(header, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			http.Error(w, "Invalid Authorization header format", http.StatusUnauthorized)
			return
		}

		tokenString := headerParts[1]
		parts := strings.Split(tokenString, ".")
		if len(parts) != 3 {
			http.Error(w, "Invalid token format", http.StatusUnauthorized)
			return
		}

		jwtHeader := parts[0]
		jwtPayload := parts[1]
		jwtSignature := parts[2]

		message := jwtHeader + "." + jwtPayload

		byteArrSecret := []byte(m.cnf.JWTSecret)
		byteArrMessage := []byte(message)

		h := hmac.New(sha256.New, byteArrSecret)
		h.Write(byteArrMessage)
		hash := h.Sum(nil)
		expectedSignature := util.Base64UrlEncode(hash)

		if !hmac.Equal([]byte(expectedSignature), []byte(jwtSignature)) {
			http.Error(w, "Invalid token signature", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
