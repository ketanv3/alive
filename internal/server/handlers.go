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
	// Collect the healthcheck results.
	results := make(map[string]interface{})
	statusCode := 200
	detailed := c.DefaultQuery("full", "false") == "true"

	for _, hc := range registry.Get() {
		result := hc.Result()
		if result.Status == "unhealthy" {
			statusCode = 503
		}

		if detailed {
			results[hc.Definition.Name] = result
		} else {
			results[hc.Definition.Name] = result.Status
		}
	}

	c.JSON(statusCode, results)
}

func getRegistryHandler(c *gin.Context) {
	c.JSON(http.StatusOK, registry.GetList())
}
