package services

import (
	"bitaksi/client"
	"bitaksi/customErrors"
	"bitaksi/models"
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

type driverMongoController struct{}

var DriverMongo driverMongoController

func (c *driverMongoController) InsertOneDriver(driver models.Driver) *customErrors.HttpError {
	coordinates, coordErr := models.NewCoordinates([2]float64{driver.Location.Coordinates[0], driver.Location.Coordinates[1]})
	if coordErr != nil {
		return &customErrors.HttpError{
			StatusCode:  http.StatusBadRequest,
			ErrorString: fmt.Sprint("invalid coordinates, ", coordErr.Error()),
		}
	}

	newDriver := models.Driver{
		Location: models.Location{
			Type:        driver.Location.Type,
			Coordinates: coordinates,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	_, err := client.BitaksiInstance.Collection.InsertOne(ctx, newDriver)
	if err != nil {
		return &customErrors.HttpError{
			StatusCode:  http.StatusInternalServerError,
			ErrorString: "unable to insert driver into mongodb collection",
		}
	}

	return nil
}

func (c *driverMongoController) InsertManyDrivers(drivers []models.Driver) *customErrors.HttpError {
	newDrivers := make([]interface{}, len(drivers))
	for i, driver := range drivers {
		coordinates, coordErr := models.NewCoordinates([2]float64{driver.Location.Coordinates[0], driver.Location.Coordinates[1]})
		if coordErr != nil {
			return &customErrors.HttpError{
				StatusCode:  http.StatusBadRequest,
				ErrorString: fmt.Sprint("invalid coordinates, ", coordErr.Error()),
			}
		}

		newDrivers[i] = models.Driver{
			Location: models.Location{
				Type:        driver.Location.Type,
				Coordinates: coordinates,
			},
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	_, err := client.BitaksiInstance.Collection.InsertMany(ctx, newDrivers)
	if err != nil {
		return &customErrors.HttpError{
			StatusCode:  http.StatusInternalServerError,
			ErrorString: "unable to insert driver into mongodb collection",
		}
	}

	return nil
}

func (c *driverMongoController) GetAllDrivers() ([]models.Driver, *customErrors.HttpError) {
	var drivers []models.Driver

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	cursor, err := client.BitaksiInstance.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, &customErrors.HttpError{
			StatusCode:  http.StatusNotFound,
			ErrorString: "No records found",
		}
	}

	err = cursor.All(ctx, &drivers)
	if err != nil {
		return nil, &customErrors.HttpError{
			StatusCode:  http.StatusInternalServerError,
			ErrorString: err.Error(),
		}
	}

	return drivers, nil
}

func (c *driverMongoController) InitializeMongoDB() {
	c.clearAllData()
	f, err := os.Open("static/Coordinates.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Println("error reading coordinates csv")
		return
	}

	drivers := make([]models.Driver, len(data))
	for i, d := range data[1:] {
		lat, err := strconv.ParseFloat(d[0], 64)
		if err != nil {
			log.Println("error reading coordinate" + d[0])
		}
		long, err := strconv.ParseFloat(d[1], 64)
		if err != nil {
			log.Println("error reading coordinate" + d[1])
		}
		drivers[i] = models.Driver{
			Location: *models.NewPoint([2]float64{lat, long}),
		}
	}

	httpErr := c.InsertManyDrivers(drivers)
	if httpErr != nil {
		log.Println("Error initializing mongodb", err.Error())
	}
}

func (c *driverMongoController) clearAllData() {
	_, err := client.BitaksiInstance.Collection.DeleteMany(context.Background(), bson.M{})
	if err != nil {
		log.Println("Unable to delete data from mongodb")
	}
}
