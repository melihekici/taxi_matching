package controllers

import (
	"encoding/json"
	"fmt"
	"matching/models"
	"matching/services"
	"net/http"
)

type matchingController struct {
}

var MatchingController = &matchingController{}

// Request body for driver finding service
type findDriversRequest struct {
	// User location
	// required: true
	Location models.Location `json:"location"`
	// Radius in meters
	// required: true
	Radius float64 `json:"radius"` // in meters
}

type findDriversResponse struct {
	Driver   models.Driver `json:"driver"`
	Distance float64       `json:"distance"`
}

// swagger:route POST /find FindDriver
// Looks for a driver in given radius around given location
// responses:
//  200:
//  400:
//  409:
//  500:

// Creates a new user
func (m *matchingController) FindDrivers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Looking for drivers")

	var req findDriversRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request. " + err.Error()))
		return
	}

	driverList, err := services.GetAllDrivers(w.Header().Get("Token"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error. " + err.Error()))
		return
	}

	if len(driverList) == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No drivers nearby."))
		return
	}

	closestDriver, distance := services.FindClosestDriverInRadius(req.Location, req.Radius, driverList)
	if closestDriver == nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No drivers nearby."))
		return
	}

	resp := findDriversResponse{Driver: *closestDriver, Distance: distance}
	responseBody, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unable to serialize drivers list. " + err.Error()))
	}

	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}
