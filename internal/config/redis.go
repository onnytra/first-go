package config

import (
	"context"

	"github.com/onnytra/first-go/internal/utils"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

func NewRedisClient(logger *logrus.Logger) *redis.Client {
	hosts := utils.GetEnv("REDIS_HOST", "127.0.0.1:6379")
	password := utils.GetEnv("REDIS_PASSWORD", "")

	client := redis.NewClient(&redis.Options{
		Addr:     hosts,
		Password: password,
		DB:       0,
	})

	ctx := context.Background()
	if pong, err := client.Ping(ctx).Result(); err != nil {
		logger.Fatal("Failed to connect to redis: ", err)
	} else {
		logger.Info("Redis connected successfully: ", pong)
	}

	return client
}
