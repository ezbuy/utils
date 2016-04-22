package middleware

import (
	"net/http"
	"strings"

	"gopkg.in/labstack/echo.v1"
)

func CORS(origins []string) func(*echo.Context) error {
	return func(c *echo.Context) error {
		origin := c.Request().Header.Get("Origin")
		// 跨域为浏览器行为, 认为浏览器一定会带 Origin 头, 其余情况为客户端请求
		if origin == "" {
			return nil
		}

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
