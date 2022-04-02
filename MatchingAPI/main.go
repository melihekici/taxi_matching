package main

import (
	"log"
	"matching/handlers"
	"matching/middleware"
	"net/http"

	openApiMiddleware "github.com/go-openapi/runtime/middleware"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/find", middleware.TokenValidationMiddleware(handlers.MatchingHandler))

	// documentation
	opts1 := openApiMiddleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := openApiMiddleware.Redoc(opts1, nil)
	mux.Handle("/docs", sh)
	mux.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	err := http.ListenAndServe(":9191", mux)
	if err != nil {
		log.Println(err)
	}

}
