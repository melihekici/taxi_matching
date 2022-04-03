package middleware

import (
	"log"
	"matching/services"
	"net/http"
	"time"
)

type CircuitBreaker struct {
	url    string
	status string
}

func (d *CircuitBreaker) ChangeStatus(status string) {
	if status == "open" || status == "closed" {
		d.status = status
	}
}

func (d *CircuitBreaker) IsHealthy() bool {
	return d.status == "closed"
}

var MongodbCircuitBreaker = CircuitBreaker{url: "", status: "closed"}

func CircuitBreakerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if !MongodbCircuitBreaker.IsHealthy() {
				w.WriteHeader(http.StatusServiceUnavailable)
				w.Write([]byte("Database is not available, please check after a while"))
				return
			}
			next.ServeHTTP(w, r)
		})
}

func init() {
	go func() {
		for {
			time.Sleep(time.Second * 5)
			if !services.DriverApiHealthCheck() {
				log.Println("Circuit breaker mongodb health check: Dead")
				MongodbCircuitBreaker.ChangeStatus("open")
				time.Sleep(time.Second * 30)
			} else {
				log.Println("Circuit breaker mongodb health check: Healthy")
				MongodbCircuitBreaker.ChangeStatus("closed")
			}
		}
	}()
}
