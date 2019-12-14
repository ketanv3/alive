package main

import (
	"codezest.in/alive/internal/registry"
	"codezest.in/alive/internal/server"
)

func main() {
	configs := []string{"configs/google-com.yaml", "configs/facebook-com.yaml"}
	registry.Initialize(configs)

	// gin.SetMode(gin.ReleaseMode)
	server.Start(":8055")
}
