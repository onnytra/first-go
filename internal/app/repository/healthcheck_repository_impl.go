package repository

import "context"

type HealthCheckRepositoryImpl struct {
}

func NewHealthCheckRepository() *HealthCheckRepositoryImpl {
	return &HealthCheckRepositoryImpl{}
}

func (r *HealthCheckRepositoryImpl) GetHealthStatus(ctx context.Context) (string, error) {
	return "OK", nil
}
