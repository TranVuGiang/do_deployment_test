package handler

import (
	"net/http"

	"github.com/TranVuGiang/digital_project_deploy/internal/service"
	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func ReadinessHandler(sv *service.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		result := sv.CheckReadiness(c.Request.Context())

		statusCode := http.StatusOK
		if result.Status == "not_ready" {
			statusCode = http.StatusServiceUnavailable
		}

		c.JSON(statusCode, result)
	}
}
