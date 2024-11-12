package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/mauriciomartinezc/real-estate-mc-auth/domain"
	"github.com/mauriciomartinezc/real-estate-mc-auth/services"
	"github.com/mauriciomartinezc/real-estate-mc-common/i18n/locales"
	"github.com/mauriciomartinezc/real-estate-mc-common/utils"
)

type PermissionHandler struct {
	permissionService services.PermissionService
}

func NewPermissionHandler(permissionService services.PermissionService) *PermissionHandler {
	return &PermissionHandler{
		permissionService: permissionService,
	}
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
