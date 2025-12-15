package provider

import (
	"github.com/google/wire"
	"github.com/onnytra/first-go/internal/app/handler"
	"github.com/onnytra/first-go/internal/app/repository"
	"github.com/onnytra/first-go/internal/app/service"
)

// HealthCheckProviderSet for all healthcheck providers
var HealthCheckProviderSet = wire.NewSet(
	repository.NewHealthCheckRepository,
	wire.Bind(new(repository.HealthCheckRepository), new(*repository.HealthCheckRepositoryImpl)),

	service.NewHealthCheckService,
	wire.Bind(new(service.HealthCheckService), new(*service.HealthCheckServiceImpl)),

	handler.NewHealthCheckHandler,
	wire.Bind(new(handler.HealthCheckHandler), new(*handler.HealthCheckHandlerImpl)),
)
