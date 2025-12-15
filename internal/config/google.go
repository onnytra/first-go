package config

import (
	"context"

	"cloud.google.com/go/storage"
	"github.com/onnytra/first-go/internal/utils"
	"github.com/sirupsen/logrus"
)

// GoogleConfig holds global Google Cloud & Maps configuration.
type GoogleConfig struct {
	MapsAPIKey    string
	ProjectID     string
	BucketName    string
	StorageAPIURI string
	Credentials   string
}

// NewGoogleConfig loads Google-related configuration from environment variables.
func NewGoogleConfig(logger *logrus.Logger) *GoogleConfig {
	cfg := &GoogleConfig{
		MapsAPIKey:    utils.GetEnv("GOOGLE_MAPS_API_KEY", ""),
		ProjectID:     utils.GetEnv("GOOGLE_CLOUD_PROJECT_ID", ""),
		BucketName:    utils.GetEnv("GOOGLE_CLOUD_STORAGE_BUCKET", ""),
		StorageAPIURI: utils.GetEnv("GOOGLE_CLOUD_STORAGE_API_URI", "https://storage.googleapis.com"),
		Credentials:   utils.GetEnv("GOOGLE_APPLICATION_CREDENTIALS", ""),
	}

	logger.Infof("Google configuration loaded: ProjectID=%s, Bucket=%s", cfg.ProjectID, cfg.BucketName)
	return cfg
}

// NewGCSClient creates and returns a Google Cloud Storage client.
func NewGCSClient(logger *logrus.Logger, cfg *GoogleConfig) *storage.Client {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		logger.Fatalf("Failed to initialize GCS client: %v", err)
	}
	logger.Infof("GCS client initialized (ProjectID=%s, Bucket=%s)", cfg.ProjectID, cfg.BucketName)
	return client
}
