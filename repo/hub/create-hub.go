package hub

import (
	"app/model"
	"app/repo"

	"github.com/yanzay/log"
)

// CreateHub :
func CreateHub(hub *model.Hub) (id string, err error) {

	if dbc := repo.GetDB().Create(hub); dbc.Error != nil {
		log.Error(dbc.Error)
		err = dbc.Error
		return
	}

	id = hub.ID
	return
}
