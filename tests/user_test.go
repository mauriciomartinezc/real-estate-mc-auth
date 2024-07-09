package tests

import (
	"github.com/labstack/echo/v4"
	"github.com/mauriciomartinezc/real-estate-mc-auth/handler"
	"github.com/mauriciomartinezc/real-estate-mc-auth/service"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRegisterUser(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(`{"name":"Test User","email":"test@example.com","password":"password"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	userService := service.NewUserService(mockUserRepo, mockRoleRepo, mockPermissionRepo)
	userHandler := handler.NewUserHandler(e, userService)

	err := userHandler.Register(c)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Contains(t, rec.Body.String, "User registered successfully")
	}
}
