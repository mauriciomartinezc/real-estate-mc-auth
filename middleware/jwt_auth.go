package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/mauriciomartinezc/real-estate-mc-auth/utils"
	"strings"
)

func JWTAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return utils.SendBadRequest(c, "Missing Authorization header")
		}
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return utils.SendBadRequest(c, "Authorization header must be in the format 'Bearer {token}'")
		}

		tokenString := parts[1]
		token, claims, err := utils.ValidateToken(tokenString)
		if err != nil || !token.Valid {
			return utils.SendBadRequest(c, "Invalid or expired JWT")
		}

		c.Set("userId", claims.User.ID)
		c.Set("user", claims.User)

		return next(c)
	}
}
