package tasks

import (
	"github.com/labstack/echo/v4"
)

type Router struct {
	controller *Controller
}

func ConfigTaskRoutes(e *echo.Echo) *echo.Echo {
	r := Router{}

	tasks := e.Group("tasks")
	tasks.POST("/create", r.controller.Create)

	return e
}
