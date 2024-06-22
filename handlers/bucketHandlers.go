package handlers

import (
	errorwrapper "flash-card-language-learning/pkg/errorWrapper"
	"net/http"

	"github.com/labstack/echo/v4"
)

func HandleGetBucketAll(c echo.Context) error {
	return errorwrapper.WriteJSON(c, http.StatusOK, "All buckets")
}

func HandleGetBucketByName(c echo.Context) error {
	bucketName := c.Param("name")
	return errorwrapper.WriteJSON(c, http.StatusOK, bucketName)
}
