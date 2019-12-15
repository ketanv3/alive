package server

import (
	"net/http"

	"codezest.in/alive/internal/registry"
	"github.com/gin-gonic/gin"
)

func pingHandler(c *gin.Context) {
	c.String(http.StatusOK, "PONG")
}

func healthCheckHandler(c *gin.Context) {
	response := make(map[string]interface{})
	statusCode := 200
	detailed := c.DefaultQuery("full", "false") == "true"

	for name, result := range runHealthCheck() {
		if result.Status == "unhealthy" {
			statusCode = 503
		}

		if detailed {
			response[name] = result
		} else {
			response[name] = result.Status
		}
	}

	c.JSON(statusCode, response)
}

func getRegistryHandler(c *gin.Context) {
	c.JSON(http.StatusOK, registry.GetList())
}
