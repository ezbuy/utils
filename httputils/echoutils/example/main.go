package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"

	"github.com/ezbuy/utils/httputils/echoutils"
	"github.com/ezbuy/utils/httputils/echoutils/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.RequestTimer())
	e.Use(middleware.JSONRecover())
	e.Use(middleware.CORS([]string{"localhost"}))
	e.SetHTTPErrorHandler(echoutils.JSONErrorHandler())

	e.Any("/panic", func(c *echo.Context) error {
		panic(fmt.Errorf("panic handler"))
	})

	e.Any("/ok", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, echoutils.OKMessage)
	})

	e.Run(":9999")
}
