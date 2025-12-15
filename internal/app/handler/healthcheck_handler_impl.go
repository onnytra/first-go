package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/onnytra/first-go/internal/app/service"
	"github.com/onnytra/first-go/internal/utils"
)

type HealthCheckHandlerImpl struct {
	HealthCheckService service.HealthCheckService
}

func NewHealthCheckHandler(healthCheckService service.HealthCheckService) *HealthCheckHandlerImpl {
	return &HealthCheckHandlerImpl{
		HealthCheckService: healthCheckService,
	}
}

func (h *HealthCheckHandlerImpl) GetHealthStatus(c *gin.Context) {
	status, err := h.HealthCheckService.GetHealthStatus(c.Request.Context())
	if err != nil {
		utils.HTTPResponse(c, http.StatusInternalServerError, utils.ErrFailedServerError, nil, err)
		return
	}

	utils.HTTPResponse(c, http.StatusOK, utils.SuccessMessage, status, nil)
}
