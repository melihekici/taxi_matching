package services

import (
	"bitaksi/client"
	"bitaksi/customErrors"
	"bitaksi/models"
	"context"
	"fmt"
	"net/http"
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
