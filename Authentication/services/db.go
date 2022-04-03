package services

import (
	"auth/client"
	"context"
	"time"
)

func PostgresDBHealthCheck() bool {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Second*5))
	defer cancel()
	err := client.PostgresConnection().PingContext(ctx)
	return err == nil
}
