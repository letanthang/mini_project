package user

import (
	"app/model"
	"app/repo"
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
)

// GetUserByEmail :
func GetUserByEmail(email string) (user model.User, err error) {
	if dbc := repo.GetDB().Where("email=?", email).Limit(1).Find(&user); dbc.Error != nil {
		fmt.Println(dbc.Error)
		err = dbc.Error
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = nil
		}
		return
	}
	return
}
