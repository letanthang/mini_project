package handler

import (
	"app/model"
	teamRepo "app/repo/team"
	"fmt"

	"github.com/labstack/echo/v4"
)

// GetTeams :
func GetTeams(c echo.Context) error {
	teams, err := teamRepo.GetTeams()
	if err != nil {
		return c.JSON(retErrorServer(err))
	}
	return c.JSON(retSuccess(teams))
}

// CreateTeam :
func CreateTeam(c echo.Context) (err error) {
	type myRequest struct {
		Name     string `json:"name" query:"name" validate:"required"`
		TeamType string `json:"team_type" query:"team_type" validate:"required"`
		HubID    string `json:"hub_id" query:"hub_id" validate:"required"`
	}
	request := new(myRequest)
	if err = c.Bind(request); err != nil {
		return c.JSON(retErrorUser(fmt.Errorf("Input error: %s", err)))
	}
	if err = c.Validate(request); err != nil {
		return c.JSON(retErrorUser(fmt.Errorf("Input validate error: %s", err)))
	}

	data := &model.Team{
		Name:     request.Name,
		TeamType: request.TeamType,
		HubID:    request.HubID,
	}

	id, err := teamRepo.CreateTeam(data)
	if err != nil || id == "" {
		return c.JSON(retErrorServer(fmt.Errorf("Input validate error: %s", err)))
	}
	return c.JSON(retSuccess(data))
}
