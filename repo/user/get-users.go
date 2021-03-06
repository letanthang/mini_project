package user

import (
	"app/model"
	"app/repo"
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
)

func GetUsers() (users []model.User, err error) {
	if dbc := repo.GetDB().Find(&users); dbc.Error != nil {
		fmt.Println(dbc.Error)
		err = dbc.Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = nil
		}
		return
	}
	return
}
