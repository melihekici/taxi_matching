package controllers

import (
	"bitaksi/client"
	"bitaksi/models"
	"bitaksi/services"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type driverController struct {
	Id       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Location models.Location    `json:"location" bson:"location"`
}

func httpError(w http.ResponseWriter, statusCode int, errorMessage string) {
	w.WriteHeader(statusCode)
	w.Write([]byte(errorMessage))
}

var DriverController = &driverController{}

// swagger:route GET /drivers GetDrivers
// Returns all drivers
// responses:
//  200:
//  404:
//  500:

// Get All drivers
func (d *driverController) GetDrivers(w http.ResponseWriter, r *http.Request) {
	drivers, httpErr := services.DriverMongo.GetAllDrivers()
	if httpErr != nil {
		httpErr.SendResponse(w)
		return
	}

	driversBytes, err := json.Marshal(drivers)
	if err != nil {
		httpError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(driversBytes)
}

// Returns one driver
func (d *driverController) GetDriver(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
}

// swagger:route POST /drivers CreateDriver
// Creates and saves a driver to mongodb
// responses:
//  201:
//  400:
//  500:

// Creates one driver
func (d *driverController) CreateDriver(w http.ResponseWriter, r *http.Request) {
	var driver models.Driver

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	err := decoder.Decode(&driver)
	if err != nil || driver.IsNil() {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}

	httpErr := services.DriverMongo.InsertOneDriver(driver)
	if httpErr != nil {
		httpErr.SendResponse(w)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Driver created succesfully")
}

// swagger:route POST /drivers/batch CreateDrivers
// Creates and saves drivers in batch to mongodb
// responses:
//  201:
//  400:
//  500:

// Create drivers in batch
func (d *driverController) CreateDrivers(w http.ResponseWriter, r *http.Request) {
	var drivers models.Drivers

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	err := decoder.Decode(&drivers)
	log.Println("-------")
	log.Println(drivers)
	log.Println("-------")
	if err != nil || drivers.IsNil() || drivers.HasNilDriver() {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request"))
		return
	}

	httpErr := services.DriverMongo.InsertManyDrivers(drivers.Drivers)
	if httpErr != nil {
		httpErr.SendResponse(w)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "Drivers are created succesfully")
}

// swagger:route DELETE /drivers/{id} DeleteDriver
// Creates and saves drivers in batch to mongodb
// responses:
//  201:
//  400:
//  500:

// Delete one driver
func (d *driverController) DeleteDriver(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/drivers/"):]

	driverId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unable to parse driver id"))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()
	res, err := client.BitaksiInstance.Collection.DeleteOne(ctx, bson.M{"_id": driverId})
	if res.DeletedCount == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "User with id: %s not found.", driverId)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error while deleting driver. ", err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successfully deleted driver."))
}

// Delete drivers in batch
func (d *driverController) DeleteDrivers(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL)
}
