package middleware

import (
	"bitaksi/services"
	"net/http"
)

func TokenValidationMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if _, ok := r.Header["Token"]; !ok {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Token is Missing"))
				return
			}

			token := r.Header["Token"][0]
			_, httpErr := services.ValidateAuthentication(token)
			if httpErr != nil {
				httpErr.SendResponse(w)
				return
			}

			next.ServeHTTP(w, r)
		})
}
