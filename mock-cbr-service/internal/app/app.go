package app

import (
	"mock-cbr-service/internal/config"
	"mock-cbr-service/internal/logger"

	"github.com/gin-gonic/gin"
)

type App struct {
	r   *gin.Engine
	cfg *config.Config
	log logger.Logger
}

func New(cfg *config.Config, log logger.Logger) *App {
	r := gin.New()

	app := &App{
		r:   r,
		cfg: cfg,
		log: log,
	}

	return app
}

func (a *App) Run() error {
	if err := a.r.Run(); err != nil {
		a.log.Error("can not start server")
		return err
	}

	return nil
}

func (a *App) Stop() {

}
