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

	appConfig := newConfig()
	if appConfig == nil {
		e.Logger.Fatal(errors.New("failed to get config"))
	}

	e.Logger.Fatal(startTaskOrchestrator(e, appConfig))
}

func startTaskOrchestrator(e *echo.Echo, appConfig *config.Config) error {
	return e.Start(appConfig.Port)
}

func registerRoutes(e *echo.Echo) *echo.Echo {
	e = tasks.ConfigTaskRoutes(e)

	return e
}

func newLogger() echo.MiddlewareFunc {
	return middleware.Logger()
}

func newConfig() *config.Config {
	return &config.Config{
		Port: ":8080",
	}
}
