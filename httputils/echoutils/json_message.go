package echoutils

import (
	"net/http"

	"gopkg.in/labstack/echo.v1"
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
