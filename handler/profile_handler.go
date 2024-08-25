package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/mauriciomartinezc/real-estate-mc-auth/domain"
	localesAuth "github.com/mauriciomartinezc/real-estate-mc-auth/i18n/locales"
	"github.com/mauriciomartinezc/real-estate-mc-auth/middleware"
	"github.com/mauriciomartinezc/real-estate-mc-auth/service"
	"github.com/mauriciomartinezc/real-estate-mc-common/i18n/locales"
	"github.com/mauriciomartinezc/real-estate-mc-common/utils"
)

type ProfileHandler struct {
	profileService service.ProfileService
}

func NewProfileHandler(e *echo.Group, profileService service.ProfileService) {
	handler := &ProfileHandler{
		profileService: profileService,
	}
	e.Use(middleware.JWTAuth)
	e.GET("/profile", handler.MeProfile)
	e.PUT("/profile/:uuid", handler.Update)
}

func (h *ProfileHandler) Update(c echo.Context) error {
	err := h.MeProfile(c)
	if err != nil {
		return err
	}
	profile := new(domain.Profile)
	if err := c.Bind(profile); err != nil {
		return utils.SendBadRequest(c, locales.ErrorPayload)
	}
	if _, err := h.profileService.Update(profile); err != nil {
		return utils.SendInternalServerError(c, err.Error())
	}
	return utils.SendSuccess(c, localesAuth.ProfileSuccess, profile)
}

func (h *ProfileHandler) MeProfile(c echo.Context) error {
	profile, err := h.profileService.MeProfile(c)
	if err != nil {
		return utils.SendBadRequest(c, err.Error())
	}
	return utils.SendSuccess(c, localesAuth.ProfileSuccess, profile)
}
