package main

import (
	"log"
	"matching/handlers"
	"matching/middleware"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/find", middleware.TokenValidationMiddleware(handlers.MatchingHandler))

	err := http.ListenAndServe(":9191", mux)
	if err != nil {
		log.Println(err)
	}

}
