package main

import (
	"flag"
	"fmt"

	"codezest.in/alive/internal/registry"
	"codezest.in/alive/internal/server"
)

func main() {
	port := flag.Int("p", 8055, "port number to start the server")
	flag.Parse()

	// Initialize the registry with the given list configs.
	registry.Initialize(flag.Args())

	// gin.SetMode(gin.ReleaseMode)
	server.Start(fmt.Sprintf(":%d", *port))
}
