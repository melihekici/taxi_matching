package services

import (
	"bitaksi/client"
	"bitaksi/config"
	"bitaksi/models"
	"context"
	"fmt"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
)

func Test_InsertDriver(t *testing.T) {
	client.ConnectDB()
	client.BitaksiInstance.Collection = client.BitaksiInstance.DB.Collection(config.MONGO["TESTCOLLECTION"])

	// Clear the test collection for testing purposes
	client.BitaksiInstance.Collection.DeleteMany(context.Background(), bson.M{})

	drivers := []models.Driver{
		{Location: models.Location{Type: "Point", Coordinates: [2]float64{180, -90}}},
		{Location: models.Location{Type: "Point", Coordinates: [2]float64{180, 90}}},
		{Location: models.Location{Type: "Point", Coordinates: [2]float64{-180, -90}}},
		{Location: models.Location{Type: "Point", Coordinates: [2]float64{-180, 90}}},
		{Location: models.Location{Type: "Point", Coordinates: [2]float64{180, 25}}},
		{Location: models.Location{Type: "Point", Coordinates: [2]float64{-180, 77}}},
		{Location: models.Location{Type: "Point", Coordinates: [2]float64{13, -90}}},
		{Location: models.Location{Type: "Point", Coordinates: [2]float64{123, 90}}},
		{Location: models.Location{Type: "Point", Coordinates: [2]float64{137, -45}}},
		{Location: models.Location{Type: "Point", Coordinates: [2]float64{-39, -25}}},
	}

	for _, d := range drivers {
		err := DriverMongo.InsertOneDriver(d)
		if err != nil {
			t.Error(err.Error())
		}
	}

	delResult, _ := client.BitaksiInstance.Collection.DeleteMany(context.Background(), bson.M{})
	if delResult.DeletedCount != 10 {
		t.Error("Not all drivers are inserted to the database")
	}
}

func Test_InsertDriver_LatErr(t *testing.T) {
	client.ConnectDB()
	client.BitaksiInstance.Collection = client.BitaksiInstance.DB.Collection(config.MONGO["TESTCOLLECTION"])

	// Clear the test collection for testing purposes
	client.BitaksiInstance.Collection.DeleteMany(context.Background(), bson.M{})

	drivers := []models.Driver{
		{Location: models.Location{Type: "Point", Coordinates: [2]float64{181, -90}}},
		{Location: models.Location{Type: "Point", Coordinates: [2]float64{250, 90}}},
		{Location: models.Location{Type: "Point", Coordinates: [2]float64{-181, -90}}},
		{Location: models.Location{Type: "Point", Coordinates: [2]float64{-345, 90}}},
		{Location: models.Location{Type: "Point", Coordinates: [2]float64{-190, 1000}}},
		{Location: models.Location{Type: "Point", Coordinates: [2]float64{1000, 2000}}},
		{Location: models.Location{Type: "Point", Coordinates: [2]float64{1305, -99}}},
		{Location: models.Location{Type: "Point", Coordinates: [2]float64{-1234, 44}}},
	}

	for _, d := range drivers {
		err := DriverMongo.InsertOneDriver(d)
		if err == nil {
			t.Error(fmt.Sprint("Latitude", d.Location.Coordinates[0], " should have failed because it is not in between [-180,180]"))
		}
	}

	delResult, _ := client.BitaksiInstance.Collection.DeleteMany(context.Background(), bson.M{})
	if delResult.DeletedCount != 0 {
		t.Error("There should not been any driver recorded to the database")
	}
}

func Test_InsertDriver_LongErr(t *testing.T) {
	client.ConnectDB()
	client.BitaksiInstance.Collection = client.BitaksiInstance.DB.Collection(config.MONGO["TESTCOLLECTION"])

	// Clear the test collection for testing purposes
	client.BitaksiInstance.Collection.DeleteMany(context.Background(), bson.M{})

	drivers := []models.Driver{
		{Location: models.Location{Type: "Point", Coordinates: [2]float64{180, -91}}},
		{Location: models.Location{Type: "Point", Coordinates: [2]float64{-180, 91}}},
		{Location: models.Location{Type: "Point", Coordinates: [2]float64{-179, -95}}},
		{Location: models.Location{Type: "Point", Coordinates: [2]float64{-44, 97}}},
		{Location: models.Location{Type: "Point", Coordinates: [2]float64{-103, 1000}}},
		{Location: models.Location{Type: "Point", Coordinates: [2]float64{100, -2000}}},
		{Location: models.Location{Type: "Point", Coordinates: [2]float64{135, -99}}},
		{Location: models.Location{Type: "Point", Coordinates: [2]float64{-134, 444}}},
	}

	for _, d := range drivers {
		err := DriverMongo.InsertOneDriver(d)
		if err == nil {
			t.Error(fmt.Sprint("Longitude", d.Location.Coordinates[1], " should have failed because it is not in between [-90,90]"))
		}
	}

	delResult, _ := client.BitaksiInstance.Collection.DeleteMany(context.Background(), bson.M{})
	if delResult.DeletedCount != 0 {
		t.Error("There should not been any driver recorded to the database")
	}
}
