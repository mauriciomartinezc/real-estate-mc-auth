package utils

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func SendResponse(c echo.Context, status int, message string, data interface{}) error {
	return c.JSON(status, Response{
		Status:  status,
		Message: message,
		Data:    data,
	})
}

func SendError(c echo.Context, status int, message string, data interface{}) error {
	return c.JSON(status, Response{
		Status:  status,
		Message: message,
		Data:    data,
	})
}

func SendSuccess(c echo.Context, message string, data interface{}) error {
	return SendResponse(c, http.StatusOK, message, data)
}

func SendCreated(c echo.Context, message string, data interface{}) error {
	return SendResponse(c, http.StatusCreated, message, data)
}

func SendBadRequest(c echo.Context, message string) error {
	return SendError(c, http.StatusBadRequest, message, nil)
}

func SendInternalServerError(c echo.Context, message string) error {
	return SendError(c, http.StatusInternalServerError, message, nil)
}
