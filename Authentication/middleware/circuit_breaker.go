package middleware

import (
	"auth/services"
	"log"
	"net/http"
	"time"
)

type CircuitBreaker struct {
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

var PostgredbCircuitBreaker = CircuitBreaker{status: "closed"}

func CircuitBreakerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			if !PostgredbCircuitBreaker.IsHealthy() {
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
			if !services.PostgresDBHealthCheck() {
				log.Println("Circuit breaker postgres health check: Dead")
				PostgredbCircuitBreaker.ChangeStatus("open")
				time.Sleep(time.Second * 30)
			} else {
				log.Println("Circuit breaker postgres health check: Healthy")
				PostgredbCircuitBreaker.ChangeStatus("closed")
			}
		}
	}()
}
