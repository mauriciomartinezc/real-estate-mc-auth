package handlers

import (
	"github.com/labstack/echo/v4"
	"github.com/mauriciomartinezc/real-estate-mc-auth/domain"
	"github.com/mauriciomartinezc/real-estate-mc-auth/services"
	utilsAuth "github.com/mauriciomartinezc/real-estate-mc-auth/utils"
	"github.com/mauriciomartinezc/real-estate-mc-common/i18n/locales"
	"github.com/mauriciomartinezc/real-estate-mc-common/utils"
)

type CompanyHandler struct {
	companyService services.CompanyService
}

func NewCompanyHandler(companyService services.CompanyService) *CompanyHandler {
	return &CompanyHandler{companyService: companyService}
}

func (h *CompanyHandler) CreateCompany(c echo.Context) error {
	company := new(domain.Company)
	if err := c.Bind(company); err != nil {
		return utils.SendBadRequest(c, locales.ErrorPayload)
	}

	if err := validate.Struct(company); err != nil {
		return utils.SendErrorValidations(c, locales.ErrorPayload, err)
	}

	user := utilsAuth.UserAuth(c)

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

	company, err := h.companyService.FindById(companyUuid)

	if err = c.Bind(company); err != nil {
		return utils.SendBadRequest(c, locales.ErrorPayload)
	}

	if err = validate.Struct(company); err != nil {
		return utils.SendErrorValidations(c, locales.ErrorPayload, err)
	}

	err = h.companyService.Update(company)
	if err != nil {
		return utils.SendInternalServerError(c, err.Error())
	}

	return utils.SendSuccess(c, locales.SuccessResponse, company)
}

func (h *CompanyHandler) CompaniesMe(c echo.Context) error {
	user := utilsAuth.UserAuth(c)
	companies, err := h.companyService.CompaniesMe(user)
	if err != nil {
		return utils.SendInternalServerError(c, err.Error())
	}
	return utils.SendSuccess(c, locales.SuccessResponse, companies)
}
