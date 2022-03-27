package main

import (
	"auth/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	mainRouter := mux.NewRouter()
	authRoter := mainRouter.PathPrefix("/auth").Subrouter()

	authRoter.HandleFunc("/signup", handlers.SignupHandler)
	authRoter.HandleFunc("/signin", handlers.SigninHandler)

	http.ListenAndServe(":9090", mainRouter)

}
