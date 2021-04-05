package controller

import (
	"github.com/bducha/img-repo/storage"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type ImageController struct{}

type PostResponse struct {
	Url string `json:"url"`
}

func (controller ImageController) Post(c echo.Context) error {

	file, err := c.FormFile("image")
	if err != nil {
		log.Fatal(err)
	}
	uid := uuid.New().String()

	store := storage.GetInstance()
	imageUrl, err := store.Store(file, uid)
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(http.StatusOK, PostResponse{Url: imageUrl})
}
