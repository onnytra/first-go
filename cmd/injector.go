//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/onnytra/first-go/internal/app/handler"
	"github.com/onnytra/first-go/internal/provider"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Dependencies struct {
	DB                 *gorm.DB
	Logger             *logrus.Logger
	HealthCheckHandler handler.HealthCheckHandler
}

func InitDependency() *Dependencies {
	wire.Build(
		provider.ConfigProviderSet,
		provider.HealthCheckProviderSet,
		wire.Struct(new(Dependencies), "*"),
	)
	return nil
}
