package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ResponseContent struct {
	Code           int         `json:"code"`
	Message        string      `json:"message"`
	Data           interface{} `json:"data"`
	DisplayMessage string      `json:"display_message,omitempty"`
}

func retSuccess(data interface{}) (int, ResponseContent) {
	return http.StatusOK, ResponseContent{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    data,
	}
}

func retErrorUser(err error) (int, ResponseContent) {
	return http.StatusBadRequest, ResponseContent{
		Code:    http.StatusBadRequest,
		Message: err.Error(),
	}
}

// ServerError :
func retErrorServer(err error) (int, ResponseContent) {
	return http.StatusBadGateway, ResponseContent{
		Code:    http.StatusBadGateway,
		Message: err.Error(),
	}
}

// NotFound func
func NotFound(err error) (int, ResponseContent) {
	return http.StatusOK, ResponseContent{
		Code:    http.StatusNotFound,
		Message: err.Error(),
	}
}

// HealthCheck :
func HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}
