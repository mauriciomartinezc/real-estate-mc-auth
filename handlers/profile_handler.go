package handlers

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/mauriciomartinezc/real-estate-mc-auth/domain"
	localesAuth "github.com/mauriciomartinezc/real-estate-mc-auth/i18n/locales"
	"github.com/mauriciomartinezc/real-estate-mc-auth/services"
	"github.com/mauriciomartinezc/real-estate-mc-common/i18n/locales"
	"github.com/mauriciomartinezc/real-estate-mc-common/utils"
)

type ProfileHandler struct {
	profileService services.ProfileService
	userService    services.UserService
}

func NewProfileHandler(profileService services.ProfileService, userService services.UserService) *ProfileHandler {
	return &ProfileHandler{
		profileService: profileService,
		userService:    userService,
	}
}

func (h *ProfileHandler) Create(c echo.Context) error {
	userId, ok := c.Get("userId").(uuid.UUID)
	if !ok {
		return utils.SendInternalServerError(c, localesAuth.CouldNotGetUserId)
	}
	user, err := h.userService.Find(userId)

	if err != nil {
		return utils.SendInternalServerError(c, err.Error())
	}

	if utils.IsValidUUID(*user.ProfileId) {
		return utils.SendSuccess(c, locales.SuccessCreated, user.Profile)
	}

	profile := new(domain.Profile)
	if err := c.Bind(profile); err != nil {
		return utils.SendBadRequest(c, locales.ErrorPayload)
	}
	if _, err := h.profileService.Create(user, profile); err != nil {
		return utils.SendInternalServerError(c, err.Error())
	}

	err = h.userService.UpdateProfileId(user, profile)
	if err != nil {
		return utils.SendInternalServerError(c, err.Error())
	}

	return utils.SendSuccess(c, locales.SuccessCreated, profile)
}

func (h *ProfileHandler) Update(c echo.Context) error {
	profile, err := h.profileService.MeProfile(c)
	if err != nil {
		return err
	}
	if err := c.Bind(profile); err != nil {
		return utils.SendBadRequest(c, locales.ErrorPayload)
	}

	if _, err := h.profileService.Update(profile); err != nil {
		return utils.SendInternalServerError(c, err.Error())
	}
	return utils.SendSuccess(c, localesAuth.ProfileSuccess, profile)
}

func (h *ProfileHandler) MeProfile(c echo.Context) error {
	_, ok := c.Get("userId").(uuid.UUID)
	if !ok {
		return utils.SendInternalServerError(c, localesAuth.CouldNotGetUserId)
	}
	profile, err := h.profileService.MeProfile(c)
	if err != nil {
		return utils.SendBadRequest(c, err.Error())
	}
	return utils.SendSuccess(c, localesAuth.ProfileSuccess, profile)
}
