package service

import (
	"context"

	"github.com/onnytra/first-go/internal/app/repository"
)

type HealthCheckServiceImpl struct {
	HealthCheckRepository repository.HealthCheckRepository
}

func NewHealthCheckService(healthCheckRepository repository.HealthCheckRepository) *HealthCheckServiceImpl {
	return &HealthCheckServiceImpl{
		HealthCheckRepository: healthCheckRepository,
	}
}

func (s *HealthCheckServiceImpl) GetHealthStatus(ctx context.Context) (string, error) {
	return s.HealthCheckRepository.GetHealthStatus(ctx)
}
