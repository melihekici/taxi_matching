package main

import (
	"log"
	"matching/handlers"
	"matching/middleware"
	"matching/services"
	"net/http"

	openApiMiddleware "github.com/go-openapi/runtime/middleware"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/find", middleware.CircuitBreakerMiddleware(middleware.TokenValidationMiddleware(handlers.MatchingHandler)))

	// documentation
	opts1 := openApiMiddleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := openApiMiddleware.Redoc(opts1, nil)
	mux.Handle("/docs", sh)
	mux.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	// driver-api-documentation
	opts2 := openApiMiddleware.RedocOpts{SpecURL: "/driver-api.yaml", Path: "driver-api-docs"}
	sh2 := openApiMiddleware.Redoc(opts2, nil)
	mux.Handle("/driver-api-docs", sh2)
	mux.HandleFunc("/driver-api.yaml", services.GetDriverApiDocs)

	err := http.ListenAndServe(":9191", mux)
	if err != nil {
		log.Println(err)
	}
}
