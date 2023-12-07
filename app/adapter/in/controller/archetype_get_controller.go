package controller

import (
	"archetype/app/shared/archetype/container"
	einar "archetype/app/shared/archetype/echo_server"

	"net/http"

	"github.com/labstack/echo/v4"
)

var _ = container.InjectInboundAdapter(func() error {
	einar.Echo.GET("/api/insert_your_pattern_here", archetypeGetController)
	return nil
})

func archetypeGetController(c echo.Context) error {
	return c.JSON(http.StatusOK, "insert_your_custom_response")
}
