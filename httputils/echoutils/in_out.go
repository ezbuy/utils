package echoutils

import (
	"encoding/json"
	"net/http"

	"github.com/labstack/echo"
)

func DecodeJSONInput(c *echo.Context, v interface{}) error {
	decoder := json.NewDecoder(c.Request().Body)

	if err := decoder.Decode(v); err != nil {
		RequestLogger(c).Errorf("invalid request data: %s", err)
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request data")
	}

	return nil
}
