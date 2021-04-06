package server

import (
	"github.com/bducha/img-repo/controller"
	"github.com/labstack/echo/v4"
)

// Declares the app routes
func declareRoutes(e *echo.Echo) {
	defaultController := controller.DefaultController{}
	imageController := controller.ImageController{}

	e.GET("/", defaultController.Home)

	e.POST("/images", imageController.Post)
	e.DELETE("/images/:uid", imageController.Delete)
}
