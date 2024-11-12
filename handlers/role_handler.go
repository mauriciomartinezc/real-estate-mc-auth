package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/mauriciomartinezc/real-estate-mc-auth/domain"
	"github.com/mauriciomartinezc/real-estate-mc-auth/services"
	"github.com/mauriciomartinezc/real-estate-mc-common/i18n/locales"
	"github.com/mauriciomartinezc/real-estate-mc-common/utils"
)

type RoleHandler struct {
	roleService services.RoleService
}

func NewRoleHandler(roleService services.RoleService) *RoleHandler {
	return &RoleHandler{
		roleService: roleService,
	}
}

func (h *RoleHandler) CreateRole(c echo.Context) error {
	role := new(domain.Role)
	if err := c.Bind(role); err != nil {
		return utils.SendBadRequest(c, locales.ErrorPayload)
	}
	if err := h.roleService.CreateRole(role); err != nil {
		return utils.SendInternalServerError(c, err.Error())
	}
	return utils.SendCreated(c, locales.SuccessCreated, role)
}
