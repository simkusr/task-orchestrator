package main

import (
	"errors"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/simkusr/task-orchestrator/config"
	"github.com/simkusr/task-orchestrator/internal/api/tasks"
)

func main() {
	e := echo.New()
	e.Use(newLogger())

	e = registerRoutes(e)

	cfg, err := newConfig()
	if cfg == nil || err != nil {
		shutDown(e, errors.New("failed to get config"))
	}

	shutDown(e, startUp(e, cfg))
}

func shutDown(e *echo.Echo, err error) {
	e.Logger.Fatal(err)
}

func startUp(e *echo.Echo, cfg *config.Config) error {

	return e.Start(cfg.Port)
}

func registerRoutes(e *echo.Echo) *echo.Echo {
	e = tasks.ConfigTaskRoutes(e)

	return e
}

func newLogger() echo.MiddlewareFunc {
	return middleware.Logger()
}

func newConfig() (*config.Config, error) {
	cfg := &config.Config{}

	err := cfg.NewConfig()

	return cfg, err
}
