package echoutils

import (
	"net/http"

	"github.com/labstack/echo"
)

func JSONErrorHandler() echo.HTTPErrorHandler {
	return func(err error, c *echo.Context) {
		if echoErr, ok := err.(*echo.HTTPError); ok {
			c.JSON(echoErr.Code(), &JSONMessage{
				Message: echoErr.Error(),
			})
			return
		}

		// others
		RequestLogger(c).Error(err.Error())
		c.JSON(http.StatusInternalServerError, InternalError)
		return
	}
}
