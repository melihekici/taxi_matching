package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"matching/algorithms"
	"matching/models"
	"net/http"
	"os"
)

func DriverApiHealthCheck() bool {
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:8080/health", os.Getenv("DRIVER_API_HOST")), nil)
	if err != nil {
		log.Println("error consturcting health check request")
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	return err == nil && resp.StatusCode == 200
}

func GetAllDrivers(token string) ([]models.Driver, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:8080/drivers", os.Getenv("DRIVER_API_HOST")), nil)
	if err != nil {
		return []models.Driver{}, err
	}

	req.Header.Set("Token", token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []models.Driver{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return []models.Driver{}, errors.New(resp.Status)
	}

	var drivers []models.Driver

	err = json.NewDecoder(resp.Body).Decode(&drivers)
	if err != nil {
		return []models.Driver{}, err
	}

	return drivers, nil
}

// Finds the drives
func FindClosestDriverInRadius(location models.Location, radius float64, drivers []models.Driver) (*models.Driver, float64) {
	minDistance := radius
	var closestDriver models.Driver

	for _, driver := range drivers {
		distance := algorithms.Haversine(location, driver.Location)
		if distance <= minDistance {
			minDistance = distance
			closestDriver = driver
		}
	}

	if closestDriver.IsNil() {
		return nil, 0
	}

	return &closestDriver, minDistance
}
