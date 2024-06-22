package handlers

import (
	errorwrapper "flash-card-language-learning/pkg/errorWrapper"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandleGetCardAll(c echo.Context) error {
	bucketName := c.Param("bucket")
	return errorwrapper.WriteJSON(c, http.StatusOK, fmt.Sprintf("bucket name: %s", bucketName))
}

func HandlerGetCardByName(c echo.Context) error {
	bucketName := c.Param("bucket")
	cardName := c.Param("name")
	return errorwrapper.WriteJSON(c, http.StatusOK, fmt.Sprintf("bucket name: %s, card name: %s", bucketName, cardName))
}
