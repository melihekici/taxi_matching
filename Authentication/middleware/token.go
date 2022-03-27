package middleware

import (
	"auth/config"
	"auth/services"
	"net/http"
)

func tokenValidationMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if _, ok := r.Header["Token"]; !ok {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Token is Missing"))
				return
			}

			token := r.Header["Token"][0]
			check, err := services.ValidateToken(token, config.SECRET)

			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Could not validate token"))
				return
			}
			if !check {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Invalid Token"))
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Authorized Token"))
		})
}
