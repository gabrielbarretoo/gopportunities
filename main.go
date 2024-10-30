package main

import (
	"github.com/gabrielbarretoo/gopportunities/config"
	"github.com/gabrielbarretoo/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("Error initializing configs: %v", err)
		return
	}

	router.Initialize()
}
