package main

import (
	"bitaksi/client"
	"bitaksi/handlers"
	"bitaksi/middleware"
	"bitaksi/services"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Connect to mongoDB client
	client.ConnectDB()
	services.DriverMongo.InitializeMongoDB()

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("asd")
	})
	mux.Handle("/drivers", middleware.TokenValidationMiddleware(handlers.DriverHandler))
	mux.Handle("/drivers/", middleware.TokenValidationMiddleware(handlers.DriverHandler))

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(err)
}
