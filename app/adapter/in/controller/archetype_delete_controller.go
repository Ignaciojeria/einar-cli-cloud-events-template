package controller

import (
	"archetype/app/shared/archetype/container"
	"archetype/app/shared/archetype/echo_server"
	"log"

	"net/http"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/labstack/echo/v4"
)

var _ = container.InjectInboundAdapter(func() error {
	echo_server.Echo.DELETE("/api/insert_your_pattern_here", archetypeDeleteController)
	return nil
})

func archetypeDeleteController(c echo.Context) error {
	event, err := cloudevents.NewEventFromHTTPRequest(c.Request())
	if err != nil {
		log.Printf("failed to parse CloudEvent from request: %v", err)
		return c.JSON(http.StatusBadRequest, "insert_your_custom_response")
	}
	return c.JSON(http.StatusOK, event.String())
}
