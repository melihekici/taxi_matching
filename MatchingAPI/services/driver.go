package services

import "net/http"

func GetAllDrivers() {
	resp, err := http.Get("http://localhost")
}
