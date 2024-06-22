package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	errorwrapper "github.com/phonghaido/golang-shared-libs/pkg/errorWrapper"
)

func HandleGetCardAll(c echo.Context) error {
	bucketName := c.Param("bucket")
	return errorwrapper.EchoWriteJSON(c, http.StatusOK, fmt.Sprintf("bucket name: %s", bucketName))
}

func HandlerGetCardByName(c echo.Context) error {
	bucketName := c.Param("bucket")
	cardName := c.Param("name")
	return errorwrapper.EchoWriteJSON(c, http.StatusOK, fmt.Sprintf("bucket name: %s, card name: %s", bucketName, cardName))
}
