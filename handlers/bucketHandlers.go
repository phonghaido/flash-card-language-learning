package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	errorwrapper "github.com/phonghaido/golang-shared-libs/pkg/errorWrapper"
)

func HandleGetBucketAll(c echo.Context) error {
	return errorwrapper.EchoWriteJSON(c, http.StatusOK, "All buckets")
}

func HandleGetBucketByName(c echo.Context) error {
	bucketName := c.Param("name")
	return errorwrapper.EchoWriteJSON(c, http.StatusOK, bucketName)
}
