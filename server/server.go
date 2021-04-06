package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Initializes the echo server
func Init() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.Static("/img", "var/images")

	declareRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
