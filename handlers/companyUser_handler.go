package handlers

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/mauriciomartinezc/real-estate-mc-auth/domain"
	"github.com/mauriciomartinezc/real-estate-mc-auth/domain/request"
	"github.com/mauriciomartinezc/real-estate-mc-auth/domain/response"
	localesAuth "github.com/mauriciomartinezc/real-estate-mc-auth/i18n/locales"
	"github.com/mauriciomartinezc/real-estate-mc-auth/services"
	utilsAuth "github.com/mauriciomartinezc/real-estate-mc-auth/utils"
	"github.com/mauriciomartinezc/real-estate-mc-common/i18n/locales"
	"github.com/mauriciomartinezc/real-estate-mc-common/utils"
)

type CompanyUserHandler struct {
	companyUserService services.CompanyUserService
	userService        services.UserService
	profileService     services.ProfileService
	companyService     services.CompanyService
}

func NewCompanyUserHandler(companyUserService services.CompanyUserService, userService services.UserService, profileService services.ProfileService, companyService services.CompanyService) *CompanyUserHandler {
	return &CompanyUserHandler{companyUserService: companyUserService, userService: userService, profileService: profileService, companyService: companyService}
}

func (h *CompanyUserHandler) AddUserToCompany(c echo.Context) error {
	var companyUserRequest = request.CompanyUserRequest{}

	if err := c.Bind(&companyUserRequest); err != nil {
		return utils.SendBadRequest(c, locales.ErrorPayload)
	}

	if err := validate.Struct(companyUserRequest); err != nil {
		return utils.SendErrorValidations(c, locales.ErrorPayload, err)
	}

	userAuth := utilsAuth.UserAuth(c)

	company, err := h.companyService.FindById(companyUserRequest.CompanyId)

	if err != nil {
		return utils.SendBadRequest(c, locales.ErrorPayload)
	}

	user, err := h.userService.Find(uuid.MustParse(companyUserRequest.UserId))

	if err != nil {
		return utils.SendBadRequest(c, locales.ErrorPayload)
	}

	if err = h.companyUserService.AddUserToCompany(userAuth, user, company, companyUserRequest.Roles); err != nil {
		if err.Error() == localesAuth.UserAlreadyAssociatedCompany {
			return utils.SendBadRequest(c, err.Error())
		}
		return utils.SendInternalServerError(c, err.Error())
	}

	return utils.SendCreated(c, locales.SuccessCreated, nil)
}

func (h *CompanyUserHandler) UpdateCompanyUser(c echo.Context) error {
	companyUserId := c.Param("uuid")

	companyUser, err := h.companyUserService.FindById(companyUserId)

	if err != nil {
		return utils.SendBadRequest(c, locales.InvalidUuid)
	}

	companyUserRequest := request.CompanyUserRequest{}

	if err = c.Bind(&companyUserRequest); err != nil {
		return utils.SendBadRequest(c, locales.ErrorPayload)
	}

	if err = validate.Struct(companyUserRequest); err != nil {
		return utils.SendErrorValidations(c, locales.ErrorPayload, err)
	}

	userAuth := utilsAuth.UserAuth(c)

	if err = h.companyUserService.SyncUserRolesInCompany(userAuth, companyUser, companyUserRequest.Roles); err != nil {
		return utils.SendInternalServerError(c, err.Error())
	}

	return utils.SendSuccess(c, locales.SuccessResponse, nil)
}

func (h *CompanyUserHandler) FindById(c echo.Context) error {
	companyUserId := c.Param("uuid")

	companyUser, err := h.companyUserService.FindById(companyUserId, "User.Profile", "Creator.Profile", "Updater.Profile", "Roles")
	if err != nil {
		return utils.SendBadRequest(c, locales.InvalidUuid)
	}

	companyUserResponse := response.ToCompanyUserResponse(*companyUser)
	return utils.SendSuccess(c, locales.SuccessResponse, companyUserResponse)
}

func (h *CompanyUserHandler) DeleteCompanyUser(c echo.Context) error {
	companyUserId := c.Param("uuid")

	companyUser, err := h.companyUserService.FindById(companyUserId)

	if err != nil {
		return utils.SendBadRequest(c, locales.InvalidUuid)
	}

	if err = h.companyUserService.Delete(companyUser); err != nil {
		return utils.SendInternalServerError(c, err.Error())
	}

	return utils.SendSuccess(c, locales.SuccessResponse, nil)
}

func (h *CompanyUserHandler) CreateUser(c echo.Context) error {
	userRequest := request.UserRequest{}

	if err := c.Bind(&userRequest); err != nil {
		return utils.SendBadRequest(c, locales.ErrorPayload)
	}

	if err := validate.Struct(userRequest); err != nil {
		return utils.SendErrorValidations(c, locales.ErrorPayload, err)
	}

	company, err := h.companyService.FindById(userRequest.CompanyId)

	if err != nil {
		return utils.SendBadRequest(c, locales.ErrorPayload)
	}

	userAuth := utilsAuth.UserAuth(c)

	userDomain := &domain.User{
		Email:                 userRequest.Email,
		Password:              userRequest.Password,
		CreateForUser:         true,
		RequiredResetPassword: true,
	}

	if err = h.userService.RegisterUser(userDomain); err != nil {
		if err.Error() == localesAuth.EmailAlreadyRegistered {
			return utils.SendBadRequest(c, err.Error())
		}
		return utils.SendInternalServerError(c, err.Error())
	}

	userDomain, _ = h.userService.Find(userDomain.ID)
	profile := userDomain.Profile

	profile.FirstName = &userRequest.FirstName
	profile.LastName = &userRequest.LastName
	profile.CityId = &userRequest.CityId

	profile, err = h.profileService.Update(profile)

	if err != nil {
		return utils.SendInternalServerError(c, err.Error())
	}

	if err = h.companyUserService.AddUserToCompany(userAuth, userDomain, company, userRequest.Roles); err != nil {
		if err.Error() == localesAuth.UserAlreadyAssociatedCompany {
			return utils.SendBadRequest(c, err.Error())
		}
		return utils.SendInternalServerError(c, err.Error())
	}

	return utils.SendCreated(c, locales.SuccessCreated, userDomain)
}
