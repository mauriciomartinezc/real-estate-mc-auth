package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/mauriciomartinezc/real-estate-mc-auth/domain"
	"github.com/mauriciomartinezc/real-estate-mc-auth/domain/request"
	localesAuth "github.com/mauriciomartinezc/real-estate-mc-auth/i18n/locales"
	"github.com/mauriciomartinezc/real-estate-mc-auth/middleware"
	"github.com/mauriciomartinezc/real-estate-mc-auth/service"
	utilsAuth "github.com/mauriciomartinezc/real-estate-mc-auth/utils"
	"github.com/mauriciomartinezc/real-estate-mc-common/i18n/locales"
	"github.com/mauriciomartinezc/real-estate-mc-common/utils"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(e *echo.Group, userService service.UserService) {
	handler := &UserHandler{
		userService: userService,
	}
	groupRoute := e.Group("/auth")
	groupRoute.POST("/register", handler.Register)
	groupRoute.POST("/login", handler.Login)

	e.Use(middleware.JWTAuth)
	e.POST("/resetPassword", handler.ResetPassword)
}

func (h *UserHandler) Register(c echo.Context) error {
	user := new(domain.User)
	if err := c.Bind(user); err != nil {
		return utils.SendBadRequest(c, locales.ErrorPayload)
	}

	if err := validate.Struct(user); err != nil {
		return utils.SendErrorValidations(c, locales.ErrorPayload, err)
	}

	if err := h.userService.RegisterUser(user); err != nil {
		if err.Error() == localesAuth.EmailAlreadyRegistered {
			return utils.SendBadRequest(c, err.Error())
		}
		return utils.SendInternalServerError(c, err.Error())
	}
	user.Password = ""
	return utils.SendCreated(c, locales.SuccessCreated, user)
}

func (h *UserHandler) Login(c echo.Context) error {
	loginRequest := domain.LoginRequest{}
	if err := c.Bind(&loginRequest); err != nil {
		return utils.SendBadRequest(c, localesAuth.InvalidEmailOrPassword)
	}

	user, token, err := h.userService.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		return utils.SendBadRequest(c, err.Error())
	}

	loginResponse := domain.LoginResponse{
		Token: token,
		User:  *user,
	}
	return utils.SendSuccess(c, localesAuth.LoginSuccessful, loginResponse)
}

func (h *UserHandler) ResetPassword(c echo.Context) error {
	userAuth := utilsAuth.UserAuth(c)

	changePasswordRequest := request.ChangePasswordRequest{}

	if err := c.Bind(&changePasswordRequest); err != nil {
		return utils.SendBadRequest(c, locales.ErrorPayload)
	}

	if err := validate.Struct(changePasswordRequest); err != nil {
		return utils.SendErrorValidations(c, locales.ErrorPayload, err)
	}

	if changePasswordRequest.UserId != userAuth.ID.String() {
		return utils.SendBadRequest(c, locales.ErrorPayload)
	}

	if !userAuth.RequiredResetPassword {
		return utils.SendBadRequest(c, localesAuth.NotRequiredResetPassword)
	}

	if err := h.userService.ResetPassword(userAuth, changePasswordRequest.OldPassword, changePasswordRequest.NewPassword); err != nil {
		if err.Error() == localesAuth.InvalidOldPassword {
			return utils.SendBadRequest(c, err.Error())
		}
		return utils.SendInternalServerError(c, err.Error())
	}

	return utils.SendSuccess(c, localesAuth.ResetPasswordSuccess, nil)
}
