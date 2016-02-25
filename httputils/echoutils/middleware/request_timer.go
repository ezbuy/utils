package middleware

import (
	"time"

	"github.com/labstack/echo"

	"github.com/ezbuy/utils/httputils/echoutils"
)

func RequestTimer() echo.MiddlewareFunc {
	return func(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c *echo.Context) error {
			start := time.Now()

			if err := h(c); err != nil {
				c.Error(err)
			}

			stop := time.Now()

			req := c.Request()
			resp := c.Response()

			echoutils.RequestLogger(c).Infof("[%d] %s %s | %s | %d", resp.Status(), req.Method, req.URL.String(), stop.Sub(start), resp.Size())
			return nil
		}
	}

}
