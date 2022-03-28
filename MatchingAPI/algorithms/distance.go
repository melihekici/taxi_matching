package algorithms

import (
	"matching/models"
	"math"
)

const (
	EARTH_RAIDUS = 6371000 //meters
)

// Haversine finds the distance between two given locations on a sphere, returns -1 when fail
func Haversine(firstLocation models.Location, secondLocation models.Location) float64 {
	lat1, long1 := degreeToRadians(firstLocation.Coordinates[0]), degreeToRadians(firstLocation.Coordinates[1])
	lat2, long2 := degreeToRadians(secondLocation.Coordinates[0]), degreeToRadians(secondLocation.Coordinates[1])

	absLongDifference := math.Abs(long1 - long2)
	centralAngle := math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(absLongDifference))

	return EARTH_RAIDUS * centralAngle
}

func degreeToRadians(d float64) float64 {
	return d * math.Pi / 180
}
