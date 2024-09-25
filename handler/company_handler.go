package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/mauriciomartinezc/real-estate-mc-auth/domain"
	"github.com/mauriciomartinezc/real-estate-mc-auth/middleware"
	"github.com/mauriciomartinezc/real-estate-mc-auth/service"
	"github.com/mauriciomartinezc/real-estate-mc-common/i18n/locales"
	"github.com/mauriciomartinezc/real-estate-mc-common/utils"
)

type CompanyHandler struct {
	companyService service.CompanyService
}

func NewCompanyHandler(e *echo.Group, companyService service.CompanyService) {
	handler := &CompanyHandler{companyService: companyService}
	e.Use(middleware.JWTAuth)

	e.POST("/companies", handler.CreateCompany)
	e.GET("/companies/:uuid", handler.FindCompany)
	e.PUT("/companies/:uuid", handler.UpdateCompany)
}

func (h *CompanyHandler) CreateCompany(c echo.Context) error {
	company := new(domain.Company)
	if err := c.Bind(company); err != nil {
		return utils.SendBadRequest(c, locales.ErrorPayload)
	}

	if err := validate.Struct(company); err != nil {
		return utils.SendErrorValidations(c, locales.ErrorPayload, err)
	}

	userInterface := c.Get("user")
	user, _ := userInterface.(domain.User)

	err := h.companyService.Create(company, user)
	if err != nil {
		return utils.SendInternalServerError(c, err.Error())
	}

	return utils.SendCreated(c, locales.SuccessCreated, company)
}

func (h *CompanyHandler) FindCompany(c echo.Context) error {
	companyUuid := c.Param("uuid")
	if !utils.IsValidUUID(companyUuid) {
		return utils.SendBadRequest(c, locales.InvalidUuid)
	}
	company, err := h.companyService.FindById(companyUuid)
	if err != nil {
		return utils.SendBadRequest(c, err.Error())
	}
	return utils.SendSuccess(c, locales.SuccessResponse, company)
}

func (h *CompanyHandler) UpdateCompany(c echo.Context) error {
	companyUuid := c.Param("uuid")
	if !utils.IsValidUUID(companyUuid) {
		return utils.SendBadRequest(c, locales.InvalidUuid)
	}

	company := new(domain.Company)
	if err := c.Bind(company); err != nil {
		return utils.SendBadRequest(c, locales.ErrorPayload)
	}

	if err := validate.Struct(company); err != nil {
		return utils.SendErrorValidations(c, locales.ErrorPayload, err)
	}

	err := h.companyService.Update(company)
	if err != nil {
		return utils.SendInternalServerError(c, err.Error())
	}

	return utils.SendSuccess(c, locales.SuccessResponse, company)
}
