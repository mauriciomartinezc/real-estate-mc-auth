package tests

import (
	"github.com/labstack/echo/v4"
	"github.com/mauriciomartinezc/real-estate-mc-auth/handlers"
	"github.com/mauriciomartinezc/real-estate-mc-auth/services"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreatePermission(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/permissions", strings.NewReader(`{"name":"Read"}`))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	permissionService := services.NewPermissionService(mockPermissionRepo)
	permissionHandler := handlers.NewPermissionHandler(e, permissionService)

	err := permissionHandler.CreatePermission(c)
	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Contains(t, rec.Body.String, "Permission created successfully")
	}
}
