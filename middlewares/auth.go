package middlewares

import (
	"github.com/labstack/echo/v4"
)

func ValidateToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		auth := c.Request().Header.Get("Authorization")
		if auth != "SecretTokenThatYouSavedInEnvFile" {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}
