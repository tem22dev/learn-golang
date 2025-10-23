package main

import (
	"learn-golang/internal/app"
	"learn-golang/internal/config"
)

func main() {
	// Init config
	cfg := config.NewConfig()

	// Init application
	application := app.NewApplication(cfg)

	// Start server
	if err := application.Run(); err != nil {
		panic(err)
	}
}
