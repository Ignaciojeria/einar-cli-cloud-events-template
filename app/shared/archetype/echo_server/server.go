package echo_server

import (
	"archetype/app/shared/archetype/container"
	"archetype/app/shared/archetype/slog"
	"archetype/app/shared/config"
	"archetype/app/shared/constants"
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var Echo *echo.Echo

func init() {
	config.Installations.EnableHTTPServer = true

	container.InjectInstallation(func() error {
		Echo = echo.New()
		Echo.Use(middleware.Logger())
		Echo.Use(middleware.Recover())
		return nil
	})

	container.InjectHTTPServer(func() error {
		setUpRenderer(EmbeddedPatterns...)
		for _, route := range Echo.Routes() {
			fmt.Printf("Method: %v, Path: %v, Name: %v\n", route.Method, route.Path, route.Name)
		}
		err := Echo.Start(":" + config.PORT.Get())
		if err != nil {
			slog.
				Logger.
				Error("error initializing application server",
					constants.ERROR, err.Error())
			return err
		}
		return nil
	})

}
