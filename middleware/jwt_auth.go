package middleware

import (
	"github.com/labstack/echo/v4"
	localesAuth "github.com/mauriciomartinezc/real-estate-mc-auth/i18n/locales"
	utilsAuth "github.com/mauriciomartinezc/real-estate-mc-auth/utils"
	"github.com/mauriciomartinezc/real-estate-mc-common/utils"
	"strings"
)

func JWTAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return utils.SendBadRequest(c, localesAuth.MissingAuthorizationHeader)
		}
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return utils.SendBadRequest(c, localesAuth.FormatAuthorizationHeaderError)
		}

		tokenString := parts[1]
		token, claims, err := utilsAuth.ValidateToken(tokenString)
		if err != nil || !token.Valid {
			return utils.SendBadRequest(c, localesAuth.InvalidExpiredJWT)
		}

		c.Set("userId", claims.User.ID)
		c.Set("user", claims.User)

		return next(c)
	}
}
