package team

import (
	"app/model"
	"app/repo"

	"github.com/yanzay/log"
)

// CreateTeam :
func CreateTeam(team *model.Team) (id string, err error) {

	if dbc := repo.GetDB().Create(team); dbc.Error != nil {
		log.Error(dbc.Error)
		err = dbc.Error
		return
	}

	id = team.ID
	return
}
