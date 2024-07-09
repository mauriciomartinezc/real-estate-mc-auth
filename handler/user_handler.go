package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/mauriciomartinezc/real-estate-mc-auth/domain"
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
	groupRoute.GET("/me", handler.Me)
}

func (h *UserHandler) Register(c echo.Context) error {
	user := new(domain.User)
	if err := c.Bind(user); err != nil {
		return utils.SendBadRequest(c, "Invalid request payload")
	}
	if err := h.userService.RegisterUser(user); err != nil {
		return utils.SendInternalServerError(c, err.Error())
	}
	user.Password = ""
	return utils.SendCreated(c, "User registered successfully", user)
}

func (h *UserHandler) Login(c echo.Context) error {
	loginRequest := domain.LoginRequest{}
	if err := c.Bind(&loginRequest); err != nil {
		return utils.SendError(c, http.StatusBadRequest, "Invalid email or password", nil)
	}

	user, token, err := h.userService.Login(loginRequest.Email, loginRequest.Password)
	if err != nil {
		return utils.SendError(c, http.StatusBadRequest, err.Error(), nil)
	}

	loginResponse := domain.LoginResponse{
		Token: token,
		User:  *user,
	}
	return utils.SendSuccess(c, "Login successful", loginResponse)
}

func (h *UserHandler) Me(c echo.Context) error {
	user, err := h.userService.GetMe(c)
	if err != nil {
		return utils.SendBadRequest(c, err.Error())
	}
	return utils.SendSuccess(c, "User me successfully", user)
}
