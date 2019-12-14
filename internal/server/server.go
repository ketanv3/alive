package server

import (
	"github.com/gin-gonic/gin"
)

// Start starts the API server at the given address
func Start(address string) {
	r := gin.Default()
	r.GET("/ping", pingHandler)
	r.GET("/health", healthCheckHandler)
	r.GET("/registry", getRegistryHandler)
	r.Run(address)
}
