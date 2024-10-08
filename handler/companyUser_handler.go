package handler

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/mauriciomartinezc/real-estate-mc-auth/domain"
	"github.com/mauriciomartinezc/real-estate-mc-auth/domain/request"
	"github.com/mauriciomartinezc/real-estate-mc-auth/domain/response"
	localesAuth "github.com/mauriciomartinezc/real-estate-mc-auth/i18n/locales"
	"github.com/mauriciomartinezc/real-estate-mc-auth/middleware"
	"github.com/mauriciomartinezc/real-estate-mc-auth/service"
	"github.com/mauriciomartinezc/real-estate-mc-common/i18n/locales"
	"github.com/mauriciomartinezc/real-estate-mc-common/utils"
)

type CompanyUserHandler struct {
	userService        service.UserService
	companyService     service.CompanyService
	companyUserService service.CompanyUserService
}

func NewCompanyUserHandler(e *echo.Group, companyUserService service.CompanyUserService, userService service.UserService, companyService service.CompanyService) {
	handler := &CompanyUserHandler{companyUserService: companyUserService, userService: userService, companyService: companyService}
	e.Use(middleware.JWTAuth)

	e.GET("/companyUsers/:uuid", handler.FindById)
	e.POST("/companyUsers", handler.AddUserToCompany)
	e.PUT("/companyUsers/:uuid", handler.UpdateCompanyUser)
	e.DELETE("/companyUsers/:uuid", handler.DeleteCompanyUser)
}

func (h *CompanyUserHandler) AddUserToCompany(c echo.Context) error {
	var companyUserRequest = request.CompanyUserRequest{}

	if err := c.Bind(&companyUserRequest); err != nil {
		return utils.SendBadRequest(c, locales.ErrorPayload)
	}

	if err := validate.Struct(companyUserRequest); err != nil {
		return utils.SendErrorValidations(c, locales.ErrorPayload, err)
	}

	userInterface := c.Get("user")
	userAuth, _ := userInterface.(domain.User)

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

	userInterface := c.Get("user")
	userAuth, _ := userInterface.(domain.User)

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
