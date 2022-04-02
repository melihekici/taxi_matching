package main

import (
	"auth/handlers"
	"net/http"

	openApiMiddleware "github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
)

func main() {
	mainRouter := mux.NewRouter()
	authRoter := mainRouter.PathPrefix("/auth").Subrouter()

	authRoter.HandleFunc("/signup", handlers.SignupHandler)
	authRoter.HandleFunc("/signin", handlers.SigninHandler)

	// documentation
	opts1 := openApiMiddleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := openApiMiddleware.Redoc(opts1, nil)
	mainRouter.Handle("/docs", sh)
	mainRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	http.ListenAndServe(":9090", mainRouter)
}
