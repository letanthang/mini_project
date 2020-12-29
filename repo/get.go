package repo

import (
	"app/model"
	"fmt"
)

func GetUsers() (*[]model.User, error) {
	var users []model.User
	if dbc := GetDB().Find(&users); dbc.Error != nil {
		fmt.Println(dbc.Error)
		return nil, dbc.Error
	}
	return &users, nil
}
