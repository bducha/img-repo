package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type DefaultController struct{}

func (controller DefaultController) Home(c echo.Context) error {
	data := struct {
		Api     string `json:"api"`
		Version string `json:"version"`
	}{
		"ImgRepo API",
		"0.0.1",
	}
	return c.JSON(http.StatusOK, data)
}
