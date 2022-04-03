package services

import (
	"fmt"
	"io/ioutil"
	"log"
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

func GetDriverApiDocs(w http.ResponseWriter, r *http.Request) {
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:8080/swagger.yaml", os.Getenv("DRIVER_API_HOST")), nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("unable to get driver api documentation"))
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("unable to get driver api documentation"))
		return
	}
	defer resp.Body.Close()

	var body []byte
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("unable to get driver api documentation"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
