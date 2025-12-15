package handler

import "github.com/gin-gonic/gin"

type HealthCheckHandler interface {
	GetHealthStatus(c *gin.Context)
}
