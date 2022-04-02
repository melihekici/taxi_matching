package main

import (
	"bitaksi/client"
	"bitaksi/handlers"
	"bitaksi/middleware"
	"bitaksi/services"
	"log"
	"net/http"

	openApiMiddleware "github.com/go-openapi/runtime/middleware"
)

func main() {
	// Connect to mongoDB client
	client.ConnectDB()
	services.DriverMongo.InitializeMongoDB()

	mux := http.NewServeMux()

	mux.Handle("/drivers", middleware.TokenValidationMiddleware(handlers.DriverHandler))
	mux.Handle("/drivers/", middleware.TokenValidationMiddleware(handlers.DriverHandler))

	// documentation
	opts1 := openApiMiddleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := openApiMiddleware.Redoc(opts1, nil)
	mux.Handle("/docs", sh)
	mux.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Println(err)
	}
}
