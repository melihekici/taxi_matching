package models

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestDriverEquals(t *testing.T) {
	var id [12]byte
	copy(id[:], "asd")
	d1 := Driver{
		Id:       primitive.ObjectID(id),
		Location: Location{Coordinates: [2]float64{11, 12}},
	}
	d2 := Driver{
		Id:       primitive.ObjectID(id),
		Location: Location{Coordinates: [2]float64{11, 12}},
	}

	if !d1.Equals(&d2) || !d2.Equals(&d1) {
		t.Error("Expected equality, got unequality")
	}
}

func TestDriverIsNil(t *testing.T) {
	d1 := Driver{}
	d2 := Driver{
		Location: Location{Coordinates: [2]float64{11, 12}},
	}

	if !d1.IsNil() {
		t.Error("Method IsNil of Driver d1 should have returned true")
	}

	if d2.IsNil() {
		t.Error("Method IsNil of Driver d2 should have returned false")
	}
}

func TestDriversIsNil(t *testing.T) {
	drivers := Drivers{}
	drivers2 := Drivers{
		Drivers: []Driver{
			{},
		},
	}
	drivers3 := Drivers{
		Drivers: []Driver{
			{Location: Location{Coordinates: [2]float64{11, 12}}},
		},
	}
	if !drivers.IsNil() {
		t.Error("Method IsNil should have returned true")
	}
	if !drivers2.IsNil() {
		t.Error("Method IsNil should have returned true")
	}
	if drivers3.IsNil() {
		t.Error("Method IsNil should have returned true")
	}
}
