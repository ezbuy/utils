package middleware

import (
	"net/http"
	"runtime/debug"

	"github.com/labstack/echo"

	"github.com/ezbuy/utils/httputils/echoutils"
)

func JSONRecover() echo.Middleware {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			defer func() {
				if err := recover(); err != nil {
					if echoErr, ok := err.(*echo.HTTPError); ok {
						c.JSON(echoErr.Code(), &echoutils.JSONMessage{
							Message: echoErr.Error(),
						})
						return
					}

					echoutils.RequestLogger(c).Errorf("PANIC: %s", err)
					debug.PrintStack()
					c.JSON(http.StatusInternalServerError, echoutils.InternalError)
					return
				}
			}()

			return h(c)
		}
	}
}
