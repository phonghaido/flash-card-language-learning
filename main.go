package main

import (
	"flash-card-language-learning/handlers"
	errorwrapper "flash-card-language-learning/pkg/errorWrapper"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	bucketG := e.Group("/api/v1/buckets")

	bucketG.GET("/", errorwrapper.EchoErrorWrapper(handlers.HandleGetBucketAll))
	bucketG.GET("/:name", errorwrapper.EchoErrorWrapper(handlers.HandleGetBucketByName))

	cardG := e.Group("/api/v1/cards")
	cardG.GET("/:bucket", errorwrapper.EchoErrorWrapper(handlers.HandleGetCardAll))
	cardG.GET("/:bucket/:name", errorwrapper.EchoErrorWrapper(handlers.HandlerGetCardByName))

	e.Logger.Fatal(e.Start(":8080"))
}
