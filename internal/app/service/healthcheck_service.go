package service

import "context"

type HealthCheckService interface {
	GetHealthStatus(ctx context.Context) (string, error)
}
