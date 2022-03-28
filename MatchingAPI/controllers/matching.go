package controllers

import (
	"encoding/json"
	"fmt"
	"matching/models"
	"net/http"
)

type matchingController struct {
}

var MatchingController = &matchingController{}

type findDriversRequest struct {
	Location models.Location `json:"location"`
	Radius   float64         `json:"radius"` // in meters
}

func (m *matchingController) FindDrivers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Looking for drivers")

	var req findDriversRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}

	var driverList []models.Driver
	

}
