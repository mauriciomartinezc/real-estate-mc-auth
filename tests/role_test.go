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

func TestCreateRole(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/roles", strings.NewReader(`{"name":"Admin"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	roleService := service.NewRoleService(mockRoleRepo)
	roleHandler := handler.NewRoleHandler(e, roleService)

	err := roleHandler.CreateRole(c)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Contains(t, rec.Body.String, "Role created successfully")
	}
}
