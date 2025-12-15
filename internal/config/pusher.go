package config

import (
	"github.com/onnytra/first-go/internal/utils"
	"github.com/pusher/pusher-http-go/v5"
	"github.com/sirupsen/logrus"
)

func NewPusherClient(logger *logrus.Logger) *pusher.Client {
	appId := utils.GetEnv("PUSHER_APP_ID", "")
	appKey := utils.GetEnv("PUSHER_APP_KEY", "")
	appSecret := utils.GetEnv("PUSHER_APP_SECRET", "")
	cluster := utils.GetEnv("PUSHER_APP_CLUSTER", "ap1")

	client := &pusher.Client{
		AppID:   appId,
		Key:     appKey,
		Secret:  appSecret,
		Cluster: cluster,
		Secure:  true,
	}

	logger.Info("Pusher client initialized successfully")
	return client
}
