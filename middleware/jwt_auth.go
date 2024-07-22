package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/mauriciomartinezc/real-estate-mc-auth/i18n/locales"
	"github.com/mauriciomartinezc/real-estate-mc-auth/utils"
	"strings"
)

func JWTAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return utils.SendBadRequest(c, locales.MissingAuthorizationHeader)
		}
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return utils.SendBadRequest(c, locales.FormatAuthorizationHeaderError)
		}

		tokenString := parts[1]
		token, claims, err := utils.ValidateToken(tokenString)
		if err != nil || !token.Valid {
			return utils.SendBadRequest(c, locales.InvalidExpiredJWT)
		}

		c.Set("userId", claims.User.ID)
		c.Set("user", claims.User)

		return next(c)
	}
}
