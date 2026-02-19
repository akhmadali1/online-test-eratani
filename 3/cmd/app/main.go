package main

import (
	"log"
	"time"

	"test-3/config"
	"test-3/internal/app"
)

func main() {
	// Configuration
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	// Timezone UTC
	time.Local = time.UTC

	// Run
	app.Run(cfg)
}
