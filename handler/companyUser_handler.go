package handler

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/mauriciomartinezc/real-estate-mc-auth/domain"
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

	e.POST("/companyUsers", handler.AddUserToCompany)
	e.PUT("/companiesUsers/:uuid", handler.UpdateCompanyUser)
	//e.DELETE("/companiesUsers/:uuid", handler.DeleteCompanyUser)
}

func (h *CompanyUserHandler) AddUserToCompany(c echo.Context) error {
	var request struct {
		CompanyId string       `json:"company_id" validate:"required"`
		UserId    string       `json:"user_id" validate:"required"`
		Roles     domain.Roles `json:"roles" validate:"required"`
	}

	if err := c.Bind(&request); err != nil {
		return utils.SendBadRequest(c, locales.ErrorPayload)
	}

	if err := validate.Struct(request); err != nil {
		return utils.SendErrorValidations(c, locales.ErrorPayload, err)
	}

	userInterface := c.Get("user")
	userAuth, _ := userInterface.(domain.User)

	company, err := h.companyService.FindById(request.CompanyId)

	if err != nil {
		return utils.SendBadRequest(c, locales.ErrorPayload)
	}

	user, err := h.userService.Find(uuid.MustParse(request.UserId))

	if err != nil {
		return utils.SendBadRequest(c, locales.ErrorPayload)
	}

	if err = h.companyUserService.AddUserToCompany(userAuth, user, company, request.Roles); err != nil {
		if err.Error() == localesAuth.UserAlreadyAssociatedCompany {
			return utils.SendBadRequest(c, err.Error())
		}
		return utils.SendInternalServerError(c, err.Error())
	}

	return utils.SendCreated(c, locales.SuccessCreated, nil)
}

func (h *CompanyUserHandler) UpdateCompanyUser(c echo.Context) error {
	return utils.SendSuccess(c, locales.SuccessResponse, nil)
}
