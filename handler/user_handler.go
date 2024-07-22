package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/mauriciomartinezc/real-estate-mc-auth/domain"
	"github.com/mauriciomartinezc/real-estate-mc-auth/i18n/locales"
	"github.com/mauriciomartinezc/real-estate-mc-auth/middleware"
	"github.com/mauriciomartinezc/real-estate-mc-auth/service"
	"github.com/mauriciomartinezc/real-estate-mc-auth/utils"
	"net/http"
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

	groupRoute.Use(middleware.JWTAuth)
	groupRoute.GET("/profile", handler.Profile)
}

func (h *UserHandler) Register(c echo.Context) error {
	user := new(domain.User)
	if err := c.Bind(user); err != nil {
		return utils.SendBadRequest(c, locales.ErrorPayload)
	}
	if err := h.userService.RegisterUser(user); err != nil {
		if err.Error() == locales.EmailAlreadyRegistered {
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
		return utils.SendBadRequest(c, locales.InvalidEmailOrPassword)
	}

	user, token, err := h.userService.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		return utils.SendError(c, http.StatusBadRequest, err.Error(), nil)
	}

	loginResponse := domain.LoginResponse{
		Token: token,
		User:  *user,
	}
	return utils.SendSuccess(c, locales.LoginSuccessful, loginResponse)
}

func (h *UserHandler) Profile(c echo.Context) error {
	user, err := h.userService.GetMe(c)
	if err != nil {
		return utils.SendBadRequest(c, err.Error())
	}
	return utils.SendSuccess(c, locales.ProfileSuccess, user)
}
