package tasks

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct{}

func (c *Controller) Create(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"hello": "world",
	})
}
