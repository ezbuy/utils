package middleware

import (
	"github.com/labstack/echo"

	"github.com/ezbuy/utils/httputils/echoutils"
)

func RequestIdentifier() func(*echo.Context) error {
	return func(c *echo.Context) error {
		echoutils.RequestLogger(c)
		return nil
	}
}
