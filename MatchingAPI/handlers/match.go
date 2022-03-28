package handlers

import (
	"matching/controllers"
	"net/http"
)

type matchingHandler struct{}

var MatchingHandler = &matchingHandler{}

func (m *matchingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	switch {
	case r.Method == http.MethodPost:
		controllers.MatchingController.FindDrivers(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not Allowed"))
	}
}
