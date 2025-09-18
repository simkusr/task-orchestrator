package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/simkusr/task-orchestrator/config"
	"github.com/simkusr/task-orchestrator/internal/api/tasks"
)

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.INFO)

	e.Use(newLogger())

	e = registerRoutes(e)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	cfg, err := newConfig()
	if cfg == nil || err != nil {
		log.Fatal("failed to load config")
	}

	go func() {
		if err := startUp(e, cfg); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
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
