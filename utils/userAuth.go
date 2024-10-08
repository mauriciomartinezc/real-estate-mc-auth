package utils

import (
	"github.com/labstack/echo/v4"
	"github.com/mauriciomartinezc/real-estate-mc-auth/domain"
)

func UserAuth(c echo.Context) domain.User {
	userInterface := c.Get("user")
	userAuth, _ := userInterface.(domain.User)

	return userAuth
}
