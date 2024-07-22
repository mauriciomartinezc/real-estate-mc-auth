package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/mauriciomartinezc/real-estate-mc-auth/domain"
	"github.com/mauriciomartinezc/real-estate-mc-auth/i18n/locales"
	"github.com/mauriciomartinezc/real-estate-mc-auth/service"
	"github.com/mauriciomartinezc/real-estate-mc-auth/utils"
)

type PermissionHandler struct {
	permissionService service.PermissionService
}

func NewPermissionHandler(e *echo.Group, permissionService service.PermissionService) {
	handler := &PermissionHandler{
		permissionService: permissionService,
	}

	//e.Use(middleware.JWTAuth)
	e.POST("/permissions", handler.CreatePermission)
}

func (h *PermissionHandler) CreatePermission(c echo.Context) error {
	permission := new(domain.Permission)
	if err := c.Bind(permission); err != nil {
		return utils.SendBadRequest(c, locales.ErrorPayload)
	}
	if err := h.permissionService.CreatePermission(permission); err != nil {
		return utils.SendInternalServerError(c, err.Error())
	}
	return utils.SendCreated(c, locales.SuccessCreated, permission)
}
