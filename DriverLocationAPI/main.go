package main

import (
	"bitaksi/client"
	"bitaksi/handlers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Connect to mongoDB client
	client.ConnectDB()

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("asd")
	})
	mux.Handle("/drivers", handlers.DriverHandler)
	mux.Handle("/drivers/", handlers.DriverHandler)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(err)
}
