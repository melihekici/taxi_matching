package models

import "errors"

type latitude float64
type longitude float64
type coordinates [2]float64

type GeoLocation struct {
	Type       string            `json:"type" bson:"type"`
	Location   Location          `json:"geometry" bson:"geometry"`
	Properties map[string]string `json:"properties" bson:"properties"`
}

type Location struct {
	Type        string      `json:"type" bson:"type"`
	Coordinates coordinates `json:"coordinates" bson:"coordinates"`
}

func NewPoint(coordinates [2]float64) (*Location, error) {
	c, err := NewCoordinates(coordinates)
	if err != nil {
		return &Location{}, err
	}

	return &Location{
		Type:        "Point",
		Coordinates: c,
	}, nil
}

func NewCoordinates(coordinates [2]float64) (coordinates, error) {
	lat, err := newLatitude(coordinates[0])
	if err != nil {
		return [2]float64{0, 0}, err
	}
	long, err := newLongitude(coordinates[1])
	if err != nil {
		return [2]float64{0, 0}, err
	}
	return [2]float64{float64(lat), float64(long)}, nil
}

func newLatitude(lat float64) (latitude, error) {
	if lat < -180 || lat > 180 {
		return 0, errors.New("latitude has to be between -180 and 180 inclusive")
	}
	return latitude(lat), nil
}

func newLongitude(long float64) (longitude, error) {
	if long < -90 || long > 90 {
		return 0, errors.New("longitude has to be between -90 and 90 inclusive")
	}
	return longitude(long), nil
}
