package services

import (
	"matching/models"
	"testing"
)

func TestFindClosestDriverInRadius(t *testing.T) {
	drivers := []models.Driver{
		{Location: models.Location{Coordinates: [2]float64{12.12, 20.15}}},
		{Location: models.Location{Coordinates: [2]float64{12.57, 19.75}}},
		{Location: models.Location{Coordinates: [2]float64{12.78, 20.86}}},
	}
	r := 100000.
	l := models.Location{Coordinates: [2]float64{12, 20}}

	driver, distance := FindClosestDriverInRadius(l, r, drivers)
	if !driver.Equals(&drivers[0]) {
		t.Error("Found wrong driver")
	}
	if distance == 0 {
		t.Error("Expected non zero distance got zero distance")
	}
}

func TestFindClosestDriverInRadius_Fail(t *testing.T) {
	drivers := []models.Driver{
		{Location: models.Location{Coordinates: [2]float64{12.12, 20.15}}},
		{Location: models.Location{Coordinates: [2]float64{12.57, 19.75}}},
		{Location: models.Location{Coordinates: [2]float64{12.78, 20.86}}},
	}
	r := 10000.
	l := models.Location{Coordinates: [2]float64{12, 20}}

	driver, distance := FindClosestDriverInRadius(l, r, drivers)
	if driver != nil {
		t.Error("Should not have found any drivers in range")
	}
	if distance != 0 {
		t.Error("Expected zero distance got non zero distance")
	}
}
