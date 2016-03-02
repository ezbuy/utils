package main

import (
	"fmt"

	"github.com/labstack/echo"

	"github.com/ezbuy/utils/httputils/echoutils"
	"github.com/ezbuy/utils/httputils/echoutils/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.RequestIdentifier())
	e.Use(middleware.RequestTimer())
	e.Use(middleware.JSONRecover())
	e.Use(middleware.CORS([]string{"localhost"}))
	e.SetHTTPErrorHandler(echoutils.JSONErrorHandler())

	e.Any("/panic", func(c *echo.Context) error {
		panic(fmt.Errorf("test panic"))
	})

	e.Run(":9999")
}
