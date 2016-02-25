package echoutils

import (
	"net/http"

	"github.com/labstack/echo"
)

var (
	InternalError = echo.NewHTTPError(http.StatusInternalServerError, "internal error")

	OKMessage = &JSONMessage{
		Message: "ok",
	}

	InternalErrorMessage = &JSONMessage{
		Message: "internal error",
	}
)

type JSONMessage struct {
	Message string `json:"message"`
}
