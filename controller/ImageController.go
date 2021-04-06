package controller

import (
	"github.com/bducha/img-repo/storage"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ImageController struct{}

// The Post's response
type PostResponse struct {
	Url string `json:"url"`
}

// A struct for simple message response
type MessageResponse struct {
	Message string `json:"message"`
}

// Add a new image
func (controller ImageController) Post(c echo.Context) error {

	file, err := c.FormFile("image")
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, MessageResponse{Message: "Error while trying to add image"})
	}
	uid := uuid.New().String()

	store := storage.GetInstance()
	imageUrl, err := store.Store(file, uid)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, MessageResponse{Message: "Error while trying to add image"})
	}
	return c.JSON(http.StatusOK, PostResponse{Url: imageUrl})
}

// Deletes an image
func (controller ImageController) Delete(c echo.Context) error {
	uid := c.Param("uid")
	store := storage.GetInstance()
	err := store.Delete(uid)
	if err != nil {
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, MessageResponse{Message: "Error while trying to remove image"})
	}
	return c.JSON(http.StatusOK, MessageResponse{Message: "Image successfully deleted"})
}
