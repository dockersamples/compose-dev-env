package main

import (
	"aquafarm-management/app"
	"aquafarm-management/app/config"
	"log"
)

func main() {
	// Load config
	cfg := config.Load()

	// Initialize app
	apps := app.New(cfg)

	// Run server
	if err := apps.Run(cfg); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
