package main

import (
	"mock-cbr-service/internal/app"
	"mock-cbr-service/internal/config"
	"mock-cbr-service/internal/logger"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.Env)

	app := app.New(cfg, log)

	app.Run()
}
