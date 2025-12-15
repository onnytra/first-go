package config

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"firebase.google.com/go/v4/storage"
	"github.com/onnytra/first-go/internal/utils"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/option"
)

// FirebaseConfig holds Firebase configuration values.
type FirebaseConfig struct {
	CredentialsPath string
	DatabaseURL     string
	Bucket          string
}

// NewFirebaseConfig loads Firebase config from environment variables.
func NewFirebaseConfig(logger *logrus.Logger) *FirebaseConfig {
	cfg := &FirebaseConfig{
		CredentialsPath: utils.GetEnv("FIREBASE_CREDENTIALS", ""),
		DatabaseURL:     utils.GetEnv("FIREBASE_DATABASE_URL", ""),
		Bucket:          utils.GetEnv("FIREBASE_BUCKET", ""),
	}

	logger.Infof("Firebase config loaded (Bucket=%s)", cfg.Bucket)
	return cfg
}

// NewFirebaseApp initializes the Firebase App instance.
func NewFirebaseApp(logger *logrus.Logger, cfg *FirebaseConfig) *firebase.App {
	opt := option.WithCredentialsFile(cfg.CredentialsPath)

	app, err := firebase.NewApp(context.Background(), &firebase.Config{
		DatabaseURL:   cfg.DatabaseURL,
		StorageBucket: cfg.Bucket,
	}, opt)
	if err != nil {
		logger.Fatalf("Failed to initialize Firebase App: %v", err)
	}

	logger.Info("Firebase App initialized successfully")
	return app
}

// NewFirebaseMessaging creates an FCM client for sending push notifications.
func NewFirebaseMessaging(logger *logrus.Logger, app *firebase.App) *messaging.Client {
	client, err := app.Messaging(context.Background())
	if err != nil {
		logger.Fatalf("Failed to initialize Firebase Messaging client: %v", err)
	}
	logger.Info("Firebase Messaging client initialized")
	return client
}

// NewFirebaseStorage creates a Firebase Storage client (wrapper over GCS).
func NewFirebaseStorage(logger *logrus.Logger, app *firebase.App) *storage.Client {
	client, err := app.Storage(context.Background())
	if err != nil {
		logger.Fatalf("Failed to initialize Firebase Storage client: %v", err)
	}
	logger.Info("Firebase Storage client initialized")
	return client
}
