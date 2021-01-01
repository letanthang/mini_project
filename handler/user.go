package handler

import (
	"app/model"
	userRepo "app/repo/user"
	"fmt"

	"github.com/labstack/echo/v4"
)

// GetUsers :
func GetUsers(c echo.Context) error {
	users, err := userRepo.GetUsers()
	if err != nil {
		return c.JSON(retErrorServer(err))
	}
	return c.JSON(retSuccess(users))
}

// CreateUser :
func CreateUser(c echo.Context) (err error) {
	type myRequest struct {
		Email  string `json:"email" query:"email" validate:"required"`
		Role   string `json:"role" query:"role" validate:"required"`
		TeamID string `json:"team_id" query:"team_id" validate:"required"`
	}
	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return c.JSON(retErrorUser(fmt.Errorf("Input error: %s", err)))
	}
	if err = c.Validate(request); err != nil {
		return c.JSON(retErrorUser(fmt.Errorf("Input validate error: %s", err)))
	}

	data := &model.User{
		Email:  request.Email,
		Role:   request.Role,
		TeamID: request.TeamID,
	}

	id, err := userRepo.CreateUser(data)
	if err != nil || id == "" {
		return c.JSON(retErrorServer(fmt.Errorf("Input validate error: %s", err)))
	}
	return c.JSON(retSuccess(data))
}
