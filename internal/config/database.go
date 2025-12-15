package config

import (
	"fmt"

	"github.com/onnytra/first-go/internal/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDatabase(logger *logrus.Logger) *gorm.DB {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.GetEnv("DB_USERNAME", ""),
		utils.GetEnv("DB_PASSWORD", ""),
		utils.GetEnv("DB_HOST", "127.0.0.1"),
		utils.GetEnv("DB_PORT", "3306"),
		utils.GetEnv("DB_DATABASE", ""),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatalf("Failed to connect to database: %v", err)
	}

	logger.Info("Successfully connected to database")
	return db
}
