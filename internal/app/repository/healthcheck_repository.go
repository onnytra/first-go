package repository

import "context"

type HealthCheckRepository interface {
	GetHealthStatus(ctx context.Context) (string, error)
}
