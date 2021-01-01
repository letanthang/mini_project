package handler

import (
	"app/model"
	hubRepo "app/repo/hub"
	"fmt"

	"github.com/labstack/echo/v4"
)

// GetHubs :
func GetHubs(c echo.Context) error {
	users, err := hubRepo.GetHubs()
	if err != nil {
		return c.JSON(retErrorServer(err))
	}
	return c.JSON(retSuccess(users))
}

// CreateHub :
func CreateHub(c echo.Context) (err error) {
	type myRequest struct {
		Name         string  `json:"name" query:"name" validate:"required"`
		LocationLat  float64 `json:"location_lat" query:"location_lat" validate:"required"`
		LocationLong float64 `json:"location_long" query:"location_long" validate:"required"`
	}
	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return c.JSON(retErrorUser(fmt.Errorf("Input error: %s", err)))
	}
	if err = c.Validate(request); err != nil {
		return c.JSON(retErrorUser(fmt.Errorf("Input validate error: %s", err)))
	}

	data := &model.Hub{
		Name:         request.Name,
		LocationLat:  request.LocationLat,
		LocationLong: request.LocationLong,
	}

	id, err := hubRepo.CreateHub(data)
	if err != nil || id == "" {
		return c.JSON(retErrorServer(fmt.Errorf("Input validate error: %s", err)))
	}
	return c.JSON(retSuccess(data))
}
