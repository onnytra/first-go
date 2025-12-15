package provider

import (
	"github.com/google/wire"
	"github.com/onnytra/first-go/internal/config"
)

// ConfigProviderSet for all configuration providers global
var ConfigProviderSet = wire.NewSet(
	config.NewLogger,
	config.NewDatabase,
	config.NewRedisClient,
	config.NewPusherClient,
	config.NewGoogleConfig,
	config.NewGCSClient,
	config.NewFirebaseConfig,
	config.NewFirebaseApp,
)
