package user

import (
	"app/model"
	"app/repo"
	"fmt"
)

func GetUsers() (*[]model.User, error) {
	var users []model.User
	if dbc := repo.GetDB().Find(&users); dbc.Error != nil {
		fmt.Println(dbc.Error)
		return nil, dbc.Error
	}
	return &users, nil
}
