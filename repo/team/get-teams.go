package team

import (
	"app/model"
	"app/repo"
	"fmt"
)

// GetTeams :
func GetTeams() (*[]model.Team, error) {
	var teams []model.Team
	if dbc := repo.GetDB().Find(&teams); dbc.Error != nil {
		fmt.Println(dbc.Error)
		return nil, dbc.Error
	}
	return &teams, nil
}
