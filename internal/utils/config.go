package utils

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var (
	envMap map[string]string
	once   sync.Once
)

// loadEnv loads .env file into envMap, only once (thread-safe)
func loadEnv() {
	envMap = map[string]string{}

	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("[WARNING] .env file not found or failed to load. Using OS environment variables.")
		return
	}

	envMap, err = godotenv.Read("../.env")
	if err != nil {
		log.Println("[WARNING] Failed to read .env file content. Using OS environment variables.")
	}
}

// getEnv fetches environment variable from envMap or OS environment
func GetEnv(key string, defaultValue string) string {
	once.Do(loadEnv)

	// check from loaded .env map
	if val, ok := envMap[key]; ok && val != "" {
		return val
	}

	// check from OS environment variable
	if val := os.Getenv(key); val != "" {
		return val
	}

	// fallback to default value
	return defaultValue
}
