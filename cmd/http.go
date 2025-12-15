package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/onnytra/first-go/internal/middleware"
	"github.com/onnytra/first-go/internal/utils"
)

func ServeHTTP() {
	dependencies := InitDependency()

	r := gin.Default()

	apiV2 := r.Group("/api/v2/framework")
	apiV2.Use(middleware.AuthMiddleware())
	apiV2.GET("/health", dependencies.HealthCheckHandler.GetHealthStatus)

	r.GET("/", func(ctx *gin.Context) {
		utils.HTTPResponse(ctx, 200, "Medicare Framework Service is running!", nil, nil)
	})

	err := r.Run(":" + utils.GetEnv("APP_PORT", "8000"))
	if err != nil {
		log.Fatal(err)
	}
}
