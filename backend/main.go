package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORS())

	e.GET("/products", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World")
	})

	e.Logger.Fatal(e.Start(":3000"))
}