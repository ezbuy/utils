package middleware

import (
	"net/http"
	"strings"

	"gopkg.in/labstack/echo.v1"
)

func CORS(origins []string) func(*echo.Context) error {
	return func(c *echo.Context) error {
		origin := c.Request().Header.Get("Origin")

		allowed := false
		for _, one := range origins {
			if strings.Index(origin, one) != -1 {
				allowed = true
				break
			}
		}

		if !allowed {
			return echo.NewHTTPError(http.StatusForbidden, "access denied")
		}

		resp := c.Response()
		resp.Header().Set("Access-Control-Allow-Origin", origin)
		resp.Header().Set("Access-Control-Allow-Credentials", "true")
		resp.Header().Set("Access-Control-Allow-Headers", "Content-Type,Ajax")

		if c.Request().Method == "OPTIONS" {
			return echo.NewHTTPError(http.StatusOK, "ok")
		}

		return nil
	}
}
