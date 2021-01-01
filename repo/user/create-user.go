package user

import (
	"app/model"
	"app/repo"

	"github.com/yanzay/log"
)

// CreateUser :
func CreateUser(user *model.User) (id string, err error) {

	if dbc := repo.GetDB().Create(user); dbc.Error != nil {
		log.Error(dbc.Error)
		err = dbc.Error
		return
	}

	id = user.ID
	return
}
