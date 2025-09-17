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

	cfg, _ := newConfig() // TODO add error handling
	if cfg == nil {
		e.Logger.Fatal(errors.New("failed to get config"))
	}

	e.Logger.Fatal(startTaskOrchestrator(e, cfg))
}

func startTaskOrchestrator(e *echo.Echo, cfg *config.Config) error {

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
