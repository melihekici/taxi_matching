package handlers

import (
	"bitaksi/controllers"
	"fmt"
	"net/http"
)

type driverHandler struct{}

var DriverHandler = &driverHandler{}

func (h *driverHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var batch bool = r.URL.Path[len("/drivers/"):] == "batch"
	fmt.Println(batch)

	switch {
	case r.Method == http.MethodGet:
		controllers.DriverController.GetDrivers(w, r)
		return
	case r.Method == http.MethodPost:
		if batch {
			controllers.DriverController.CreateDrivers(w, r)
		} else {
			controllers.DriverController.CreateDriver(w, r)
		}
		return
	case r.Method == http.MethodDelete:
		controllers.DriverController.DeleteDriver(w, r)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not Allowed"))
	}
}
